package worker

import (
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

func (worker Worker) processBlock(block node.Block) {
	for _, tx := range block.Transactions {
		worker.processTransaction(tx)
	}
	for _, receipt := range block.Receipts {
		worker.processReceipt(receipt)
	}
}
