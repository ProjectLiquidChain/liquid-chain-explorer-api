package surf

import (
	"path"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/storage"
)

// Service represent service that surf provides
type Service struct {
	db             database.Database
	nodeAPI        node.API
	txStorage      storage.Storage
	blockStorage   storage.Storage
	receiptStorage storage.Storage
}

// New returns new instance of Service
func New(dbURL, nodeURL, storagePath string) Service {
	return Service{
		db:             database.New(dbURL),
		nodeAPI:        node.New(nodeURL),
		txStorage:      storage.New(path.Join(storagePath, storage.TxStoragePath)),
		blockStorage:   storage.New(path.Join(storagePath, storage.BlockStoragePath)),
		receiptStorage: storage.New(path.Join(storagePath, storage.ReceiptStoragePath)),
	}
}
