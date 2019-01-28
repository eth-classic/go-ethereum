package miner

import (
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/core"
	"github.com/openether/ethcore/core/state"
	"github.com/openether/ethcore/core/types"
	"github.com/openether/ethcore/eth/downloader"
	"github.com/openether/ethcore/event"
	"github.com/openether/ethcore/logger"
	"github.com/openether/ethcore/logger/glog"
	"github.com/openether/ethcore/pow"
)

// HeaderExtra is a freeform description.
var HeaderExtra []byte

type Miner struct {
	mux *event.TypeMux

	worker *worker

	MinAcceptedGasPrice *big.Int

	threads  int
	coinbase common.Address
	mining   int32
	eth      core.Backend
	pow      pow.PoW

	canStart    int32 // can start indicates whether we can start the mining operation
	shouldStart int32 // should start indicates whether we should start after sync
}

func New(eth core.Backend, config *core.ChainConfig, mux *event.TypeMux, pow pow.PoW) *Miner {
	miner := &Miner{eth: eth, mux: mux, pow: pow, worker: newWorker(config, common.Address{}, eth), canStart: 1}
	go miner.update()

	return miner
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (self *Miner) update() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
out:
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&self.canStart, 0)
			if self.Mining() {
				self.Stop()
				atomic.StoreInt32(&self.shouldStart, 1)
				glog.V(logger.Info).Infoln("Mining operation aborted due to sync operation")
			}
		case downloader.DoneEvent, downloader.FailedEvent:
			shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

			atomic.StoreInt32(&self.canStart, 1)
			atomic.StoreInt32(&self.shouldStart, 0)
			if shouldStart {
				self.Start(self.coinbase, self.threads)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}

func (m *Miner) SetGasPrice(price *big.Int) error {

	if price == nil {
		return nil
	}

	if m.MinAcceptedGasPrice != nil && price.Cmp(m.MinAcceptedGasPrice) == -1 {
		priceTooLowError := errors.New("Gas price lower than minimum allowed.")
		return priceTooLowError
	}

	m.worker.setGasPrice(price)

	return nil
}

func (self *Miner) Start(coinbase common.Address, threads int) {
	atomic.StoreInt32(&self.shouldStart, 1)
	self.threads = threads
	self.worker.coinbase = coinbase
	self.coinbase = coinbase

	if atomic.LoadInt32(&self.canStart) == 0 {
		glog.V(logger.Info).Infoln("Can not start mining operation due to network sync (starts when finished)")
		return
	}

	atomic.StoreInt32(&self.mining, 1)

	for i := 0; i < threads; i++ {
		self.worker.register(NewCpuAgent(i, self.pow))
	}

	mlogMinerStart.AssignDetails(
		coinbase.Hex(),
		threads,
	).Send(mlogMiner)
	glog.V(logger.Info).Infof("Starting mining operation (CPU=%d TOT=%d)\n", threads, len(self.worker.agents))

	self.worker.start()

	self.worker.commitNewWork()
}

func (self *Miner) Stop() {
	self.worker.stop()
	atomic.StoreInt32(&self.mining, 0)
	atomic.StoreInt32(&self.shouldStart, 0)
	if logger.MlogEnabled() {
		mlogMinerStop.AssignDetails(
			self.coinbase.Hex(),
			self.threads,
		).Send(mlogMiner)
	}
}

func (self *Miner) Register(agent Agent) {
	if self.Mining() {
		agent.Start()
	}
	self.worker.register(agent)
}

func (self *Miner) Unregister(agent Agent) {
	self.worker.unregister(agent)
}

func (self *Miner) Mining() bool {
	return atomic.LoadInt32(&self.mining) > 0
}

func (self *Miner) HashRate() (tot int64) {
	tot += self.pow.GetHashrate()
	// do we care this might race? is it worth we're rewriting some
	// aspects of the worker/locking up agents so we can get an accurate
	// hashrate?
	for agent := range self.worker.agents {
		tot += agent.GetHashRate()
	}
	return
}

// Pending returns the currently pending block and associated state.
func (self *Miner) Pending() (*types.Block, *state.StateDB) {
	return self.worker.pending()
}

func (self *Miner) SetEtherbase(addr common.Address) {
	self.coinbase = addr
	self.worker.setEtherbase(addr)
}
