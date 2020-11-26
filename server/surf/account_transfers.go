package surf

import (
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetAccountTransfersParams is params to GetAccount transaction
type GetAccountTransfersParams struct {
	Address string `json:"address"`
}

// GetAccountTransfersResult is result of GetAccount
type GetAccountTransfersResult struct {
	Transfers []database.Transfer `json:"transfers"`
}

// GetAccountTransfers lookup transfers for an account
func (service Service) GetAccountTransfers(r *http.Request, params *GetAccountTransfersParams, result *GetAccountTransfersResult) error {
	var account database.Account
	if err := service.db.Where(database.Account{Address: params.Address}).First(&account).Error; err != nil {
		return err
	}

	var transfers []database.Transfer
	if err := service.db.
		Joins("Transaction").
		Joins("FromAccount").
		Joins("ToAccount").
		Joins("Token").
		Where(database.Transfer{FromAccountID: account.ID}).
		Or(database.Transfer{ToAccountID: account.ID}).Error; err != nil {
		return err
	}

	result.Transfers = transfers
	return nil
}
