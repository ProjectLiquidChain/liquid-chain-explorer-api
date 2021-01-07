package surf

import (
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetTokensParams is params to GetAccount transaction
type GetTokensParams struct {
	paginationParams
}

// GetTokensResult is result of GetAccount
type GetTokensResult struct {
	paginationResult
	Tokens []database.Token `json:"tokens"`
}

// GetTokens lookup txs for an account
func (service Service) GetTokens(r *http.Request, params *GetTokensParams, result *GetTokensResult) error {
	limit := params.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	var tokens []database.Token
	if err := service.db.
		Order("tokens.id ASC").
		Offset(limit * params.Page).
		Limit(limit).
		Find(&tokens).Error; err != nil {
		return err
	}
	var count int64
	if err := service.db.
		Model(&database.Token{}).
		Count(&count).Error; err != nil {
		return err
	}

	result.Tokens = tokens
	result.TotalPages = int(math.Ceil(float64(count) / float64(limit)))

	return nil
}
