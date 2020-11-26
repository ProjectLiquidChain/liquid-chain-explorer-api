package worker

import (
	"encoding/hex"
	"encoding/json"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

func (worker Worker) storeTx(tx node.Transaction) {
	hashByte, err := hex.DecodeString(tx.Hash)
	if err != nil {
		panic(err)
	}

	txByte, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}

	if err := worker.txStorage.Update(hashByte, txByte); err != nil {
		panic(err)
	}
}

func (worker Worker) processTransaction(tx node.Transaction) error {
	sender := database.Account{Address: tx.Sender}
	if err := worker.db.
		Where(database.Account{Address: tx.Sender}).
		FirstOrInit(&sender).Error; err != nil {
		return err
	}

	receiver := database.Account{Address: tx.Receiver}
	if err := worker.db.
		Where(database.Account{Address: tx.Receiver}).
		FirstOrInit(&receiver).Error; err != nil {
		return err
	}

	transaction := database.Transaction{
		Block:      tx.BlockHeight,
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Hash:       tx.Hash,
	}
	if err := worker.db.
		Where(database.Transaction{Hash: tx.Hash}).
		FirstOrCreate(&transaction).Error; err != nil {
		return err
	}

	worker.storeTx(tx)
	return nil
}
