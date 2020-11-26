package worker

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

const transferEvent = "Transfer"

func isTransferEvent(event node.Call) bool {
	if event.Name != transferEvent {
		return false
	}

	if len(event.Args) != 4 {
		return false
	}

	return event.Args[0].Type == "address" &&
		event.Args[1].Type == "address" &&
		event.Args[2].Type == "uint64" &&
		event.Args[3].Type == "uint64"
}

func (worker Worker) storeReceipt(receipt node.Receipt) {
	hashByte, err := hex.DecodeString(receipt.Transaction)
	if err != nil {
		panic(err)
	}

	receiptByte, err := json.Marshal(receipt)
	if err != nil {
		panic(err)
	}

	if err := worker.txStorage.Update(hashByte, receiptByte); err != nil {
		panic(err)
	}
}

func (worker Worker) processReceipt(receipt node.Receipt) error {
	var tx database.Transaction
	if err := worker.db.
		Where(database.Transaction{Hash: receipt.Transaction}).
		First(&tx).Error; err != nil {
		return err
	}

	for index, event := range receipt.Events {
		if isTransferEvent(event) {
			var token database.Token

			if err := worker.db.Where(database.Token{
				Address: event.Contract,
			}).First(&token); err != nil {
				log.Printf("Unsupported Transfer event for contract %s\n", event.Contract)
				continue
			}

			from := database.Account{Address: event.Args[0].Value}
			if err := worker.db.
				Where(database.Account{Address: event.Args[0].Value}).
				FirstOrCreate(&from).Error; err != nil {
				return err
			}

			to := database.Account{Address: event.Args[1].Value}
			if err := worker.db.
				Where(database.Account{Address: event.Args[1].Value}).
				FirstOrCreate(&to).Error; err != nil {
				return err
			}

			transfer := database.Transfer{
				EventIndex:    index,
				TransactionID: tx.ID,
				TokenID:       token.ID,
				FromAccountID: from.ID,
				ToAccountID:   to.ID,
				Amount:        event.Args[2].Value,
				Memo:          event.Args[3].Value,
			}

			worker.db.Where(database.Transfer{
				TransactionID: tx.ID,
				EventIndex:    index,
			}).FirstOrCreate(&transfer)
		}
	}

	worker.storeReceipt(receipt)
	return nil
}
