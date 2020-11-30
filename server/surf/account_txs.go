package surf

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

// GetAccountTxsParams is params to GetAccount transaction
type GetAccountTxsParams struct {
	Address string `json:"address"`
}

// GetAccountTxsResult is result of GetAccount
type GetAccountTxsResult struct {
	Transactions []node.Transaction `json:"transactions"`
	Receipts     []node.Receipt     `json:"receipts"`
}

// GetAccountTxs lookup txs for an account
func (service Service) GetAccountTxs(r *http.Request, params *GetAccountTxsParams, result *GetAccountTxsResult) error {
	var account database.Account
	if err := service.db.Where(database.Account{Address: params.Address}).First(&account).Error; err != nil {
		return err
	}

	var txs []database.Transaction
	if err := service.db.
		Order("transactions.id DESC").
		Where(database.Transaction{SenderID: account.ID}).
		Or(database.Transaction{ReceiverID: account.ID}).Find(&txs).Error; err != nil {
		return err
	}

	nodeTxs := []node.Transaction{}
	receipts := []node.Receipt{}

	for _, tx := range txs {
		hashByte, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return err
		}

		// Get tx
		nodeTxByte, err := service.txStorage.Get(hashByte)
		if err != nil {
			return err
		}

		var nodeTx node.Transaction
		if err := json.Unmarshal(nodeTxByte, &nodeTx); err != nil {
			return err
		}

		nodeTxs = append(nodeTxs, nodeTx)

		// Get receipt
		receiptByte, err := service.receiptStorage.Get(hashByte)
		if err != nil {
			return err
		}

		var receipt node.Receipt
		if err := json.Unmarshal(receiptByte, &receipt); err != nil {
			return err
		}

		receipts = append(receipts, receipt)
	}

	result.Transactions = nodeTxs
	result.Receipts = receipts
	return nil
}
