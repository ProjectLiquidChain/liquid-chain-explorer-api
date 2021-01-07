package surf

import (
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetTransfersParams is params to GetAccount transaction
type GetTransfersParams struct {
	paginationParams
}

// GetTransfersResult is result of GetAccount
type GetTransfersResult struct {
	paginationResult
	Transfers []database.Transfer `json:"transfers"`
}

// GetTransfers lookup transfers for an account
func (service Service) GetTransfers(r *http.Request, params *GetTransfersParams, result *GetTransfersResult) error {
	limit := params.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	var transfers []database.Transfer
	if err := service.db.
		Joins("Token").
		Joins("ToAccount").
		Joins("FromAccount").
		Joins("Transaction").
		Order("transfers.id DESC").
		Limit(limit).
		Offset(limit * params.Page).
		Find(&transfers).Error; err != nil {
		return err
	}

	var count int64
	if err := service.db.
		Model(&database.Transfer{}).
		Joins("ToAccount").
		Joins("FromAccount").
		Count(&count).Error; err != nil {
		return err
	}

	result.Transfers = transfers
	result.TotalPages = int(math.Ceil(float64(count) / float64(limit)))

	return nil
}
