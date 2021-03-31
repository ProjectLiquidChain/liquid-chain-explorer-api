package surf

import (
	"fmt"
	"net/http"
	"time"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetTxsCountParams is params to GetAccount transaction
type GetTxsCountParams struct {
	Range string `json:"range"`
	Size  int    `json:"size"`
}

// GetTxsCountResult is result of GetAccount
type GetTxsCountResult struct {
	Counts []int `json:"counts"`
}

// GetTxsCount lookup txs for an account
func (service Service) GetTxsCount(r *http.Request, params *GetTxsCountParams, result *GetTxsCountResult) error {
	toTime := time.Now()
	counts := make([]int, params.Size)
	duration, err := time.ParseDuration(params.Range)
	if err != nil {
		return err
	}
	fmt.Println(params)
	for i := 0; i < params.Size; i++ {
		fromTime := toTime.Add(-duration)
		fmt.Println(fromTime)
		var count int64
		if err := service.db.Model(&database.Transaction{}).
			Joins("JOIN blocks ON transactions.block = blocks.height").
			Where("blocks.time > ?", fromTime.Unix()).
			Where("blocks.time <= ?", toTime.Unix()).
			Count(&count).Error; err != nil {
			return err
		}
		toTime = fromTime
		counts[i] = int(count)
	}

	result.Counts = counts
	return nil
}
