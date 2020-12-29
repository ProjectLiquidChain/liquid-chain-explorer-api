package surf

import (
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

// GetBlocksParams is params to GetAccount transaction
type GetBlocksParams struct {
	paginationParams
}

// GetBlocksResult is result of GetAccount
type GetBlocksResult struct {
	paginationResult
	Blocks []node.Block `json:"blocks"`
}

// GetBlocks lookup txs for an account
func (service Service) GetBlocks(r *http.Request, params *GetBlocksParams, result *GetBlocksResult) error {
	limit := params.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	var blocks []database.Block
	if err := service.db.
		Order("blocks.height DESC").
		Offset(limit * params.Page).
		Limit(limit).
		Find(&blocks).Error; err != nil {
		return err
	}

	var count int64
	if err := service.db.
		Model(&database.Block{}).
		Count(&count).Error; err != nil {
		return err
	}
	result.TotalPages = int(math.Ceil(float64(count) / float64(limit)))

	return nil
}
