package worker

import (
	"log"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

func (worker Worker) processBlock(block node.Block) {
	log.Println("Process block", block.Height)

	for _, tx := range block.Transactions {
		if err := worker.processTransaction(tx); err != nil {
			panic(err)
		}
	}

	for _, receipt := range block.Receipts {
		if err := worker.processReceipt(receipt); err != nil {
			panic(err)
		}
	}

	worker.db.Create(&database.Block{
		Height: uint(block.Height),
		Hash:   block.Hash,
	})
}
