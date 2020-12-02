package surf

import (
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetAccountTransfersParams is params to GetAccount transaction
type GetAccountTransfersParams struct {
	paginationParams
	Address string `json:"address"`
}

// GetAccountTransfersResult is result of GetAccount
type GetAccountTransfersResult struct {
	paginationResult
	Transfers []database.Transfer `json:"transfers"`
}

// GetAccountTransfers lookup transfers for an account
func (service Service) GetAccountTransfers(r *http.Request, params *GetAccountTransfersParams, result *GetAccountTransfersResult) error {
	var account database.Account
	if err := service.db.Where(database.Account{Address: params.Address}).First(&account).Error; err != nil {
		return err
	}

	limit := params.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	query := service.db.
		Joins("Token").
		Joins("ToAccount").
		Joins("FromAccount").
		Joins("Transaction").
		Order("transfers.id DESC").
		Where(database.Transfer{FromAccountID: account.ID}).
		Or(database.Transfer{ToAccountID: account.ID})

	var transfers []database.Transfer
	if err := query.
		Limit(limit).
		Offset(params.Offset).
		Find(&transfers).Error; err != nil {
		return err
	}
	result.Transfers = transfers

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	result.TotalPage = int(math.Ceil(float64(count) / float64(limit)))

	return nil
}
