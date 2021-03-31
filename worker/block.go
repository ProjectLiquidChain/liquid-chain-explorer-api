package worker

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

func (worker Worker) storeBlock(block node.Block) {
	hashByte, err := hex.DecodeString(block.Hash)
	if err != nil {
		panic(err)
	}

	blockByte, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}

	if err := worker.blockStorage.Update(hashByte, blockByte); err != nil {
		panic(err)
	}
}

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
		Time:   block.Time,
	})

	worker.storeBlock(block)
}
