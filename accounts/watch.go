// +build darwin,!ios freebsd linux,!arm64 netbsd solaris windows

package accounts

import (
	"time"

	"github.com/ether-core/go-ethereum/logger"
	"github.com/ether-core/go-ethereum/logger/glog"
	// FUck no
	"github.com/rjeczalik/notify"
)

type watcher struct {
	ac       caching
	starting bool
	running  bool
	ev       chan notify.EventInfo
	//evs      []notify.EventInfo
	quit chan struct{}
}

func newWatcher(ac caching) *watcher {
	return &watcher{
		ac:   ac,
		ev:   make(chan notify.EventInfo, 10),
		quit: make(chan struct{}),
	}
}

// starts the watcher loop in the background.
// Start a watcher in the background if that's not already in progress.
// The caller must hold w.ac.mu.
func (w *watcher) start() {
	if w.starting || w.running {
		return
	}
	w.starting = true
	go w.loop()
}

func (w *watcher) close() {
	close(w.quit)
}

func (w *watcher) loop() {
	defer func() {
		w.ac.muLock()
		w.running = false
		w.starting = false
		w.ac.muUnlock()
	}()

	err := notify.Watch(w.ac.getKeydir(), w.ev, notify.All)
	if err != nil {
		glog.V(logger.Warn).Errorf("can't watch %s: %v", w.ac.getKeydir(), err)
		return
	}
	defer notify.Stop(w.ev)
	glog.V(logger.Detail).Infof("now watching %s", w.ac.getKeydir())
	defer glog.V(logger.Detail).Infof("no longer watching %s", w.ac.getKeydir())

	w.ac.muLock()
	w.running = true
	w.ac.muUnlock()

	// Wait for file system events and reload.
	// When an event occurs, the reload call is delayed a bit so that
	// multiple events arriving quickly only cause a single reload.
	var (
		debounce          = time.NewTimer(0)
		debounceDuration  = 500 * time.Millisecond
		inCycle, hadEvent bool
	)
	defer debounce.Stop()
	for {
		select {
		case <-w.quit:
			return
		case <-w.ev:
			if !inCycle {
				debounce.Reset(debounceDuration)
				inCycle = true
			} else {
				hadEvent = true
			}
			//w.evs = append(w.evs, <-w.ev)
		case <-debounce.C:
			w.ac.muLock()
			//w.evs = w.ac.reload(w.evs)
			w.ac.reload()
			w.ac.muUnlock()
			if hadEvent {
				debounce.Reset(debounceDuration)
				inCycle, hadEvent = true, false

			} else {
				inCycle, hadEvent = false, false
			}
		}
	}
}
