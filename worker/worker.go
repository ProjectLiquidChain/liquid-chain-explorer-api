package worker

import (
	"time"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/storage"
	"gorm.io/gorm"
)

// Worker checks network and db balance
type Worker struct {
	db             database.Database
	nodeAPI        node.API
	txStorage      storage.Storage
	blockStorage   storage.Storage
	receiptStorage storage.Storage
}

// New returns new instance of Worker
func New(dbURL, nodeURL string, txStorage, blockStorage, receiptStorage storage.Storage) Worker {
	return Worker{
		db:             database.New(dbURL),
		nodeAPI:        node.New(nodeURL),
		txStorage:      txStorage,
		blockStorage:   blockStorage,
		receiptStorage: receiptStorage,
	}
}

// Start runs the monitor
func (worker Worker) Start() {
	for {
		latestBlock, err := worker.nodeAPI.GetLatestBlock()
		if err != nil {
			panic(err)
		}

		var block database.Block
		if err := worker.db.Order("height DESC").Limit(1).First(&block).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				panic(err)
			}
		}

		for i := uint64(block.Height + 1); i < latestBlock.Height; i++ {
			block, err := worker.nodeAPI.GetBlock(i)
			if err != nil {
				panic(err)
			}

			worker.processBlock(block)
		}

		time.Sleep(2 * time.Second)
	}
}
