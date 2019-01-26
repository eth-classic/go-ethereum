package downloader

import (
	"github.com/ether-core/go-ethereum/core"
)

type InsertChainEvent struct {
	core.ChainInsertEvent
}

type InsertReceiptChainEvent struct {
	core.ReceiptChainInsertEvent
	Pivot bool
}

type InsertHeaderChainEvent struct {
	core.HeaderChainInsertEvent
}
