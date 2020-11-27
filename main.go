package main

import (
	"os"
	"path"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/server"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/storage"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/worker"
)

func main() {
	storagePath := os.Getenv("STORAGE_PATH")
	txStorage := storage.New(path.Join(storagePath, storage.TxStoragePath))
	blockStorage := storage.New(path.Join(storagePath, storage.BlockStoragePath))
	receiptStorage := storage.New(path.Join(storagePath, storage.ReceiptStoragePath))

	w := worker.New(
		os.Getenv("DATABASE_URL"),
		os.Getenv("NODE_URL"),
		txStorage, blockStorage, receiptStorage,
	)
	go w.Start()

	s := server.New(":5556",
		os.Getenv("DATABASE_URL"),
		os.Getenv("NODE_URL"),
		txStorage, blockStorage, receiptStorage)
	s.Serve()
}
