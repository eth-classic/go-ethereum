package p2p

import (
	"net"

	"github.com/ether-core/go-ethereum/metrics"
)

// meteredConn wraps a network TCP connection for metrics.
type meteredConn struct {
	net.Conn
	markBytes func(int64)
}

func newMeteredConn(conn net.Conn, ingress bool) net.Conn {
	if ingress {
		metrics.P2PIn.Mark(1)
		return &meteredConn{conn, metrics.P2PInBytes.Mark}
	} else {
		metrics.P2POut.Mark(1)
		return &meteredConn{conn, metrics.P2POutBytes.Mark}
	}
}

func (c *meteredConn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)
	c.markBytes(int64(n))
	return
}

func (c *meteredConn) Write(b []byte) (n int, err error) {
	n, err = c.Conn.Write(b)
	c.markBytes(int64(n))
	return
}
