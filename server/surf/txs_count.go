package surf

import (
	"net/http"
	"time"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetTxsCountParams is params to GetAccount transaction
type GetTxsCountParams struct {
	Range time.Duration `json:"range"`
	Size  int           `json:"count"`
}

// GetTxsCountResult is result of GetAccount
type GetTxsCountResult struct {
	paginationResult
	Counts []int `json:"counts"`
}

// GetTxsCount lookup txs for an account
func (service Service) GetTxsCount(r *http.Request, params *GetTxsCountParams, result *GetTxsCountResult) error {
	toTime := time.Now()
	counts := make([]int, params.Size)
	for i := 0; i < params.Size; i++ {
		fromTime := toTime.Add(-params.Range)
		var count int64
		if err := service.db.Model(&database.Transaction{}).
			Joins("Block").
			Where("Block.time > ?", fromTime).
			Where("Block.time <= ?", toTime).
			Count(&count).Error; err != nil {
			return err
		}
		toTime = fromTime
		counts[i] = int(count)
	}

	result.Counts = counts
	return nil
}
