package worker

import (
	"path"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/storage"
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
func New(dbURL, nodeURL, storagePath string) Worker {
	return Worker{
		db:             database.New(dbURL),
		nodeAPI:        node.New(nodeURL),
		txStorage:      storage.New(path.Join(storagePath, storage.TxStoragePath)),
		blockStorage:   storage.New(path.Join(storagePath, storage.BlockStoragePath)),
		receiptStorage: storage.New(path.Join(storagePath, storage.ReceiptStoragePath)),
	}
}
