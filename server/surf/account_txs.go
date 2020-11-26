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
}

// GetAccountTxs lookup txs for an account
func (service Service) GetAccountTxs(r *http.Request, params *GetAccountTxsParams, result *GetAccountTxsResult) error {
	var account database.Account
	if err := service.db.Where(database.Account{Address: params.Address}).First(&account).Error; err != nil {
		return err
	}

	var txs []database.Transaction
	if err := service.db.
		Where(database.Transaction{SenderID: account.ID}).
		Or(database.Transaction{ReceiverID: account.ID}).Find(&txs).Error; err != nil {
		return err
	}

	nodeTxs := []node.Transaction{}
	for _, tx := range txs {
		hashByte, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return err
		}

		nodeTxByte, err := service.txStorage.Get(hashByte)
		if err != nil {
			return err
		}

		var nodeTx node.Transaction
		if err := json.Unmarshal(nodeTxByte, &nodeTx); err != nil {
			return err
		}

		nodeTxs = append(nodeTxs, nodeTx)
	}

	result.Transactions = nodeTxs
	return nil
}
