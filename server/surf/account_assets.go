package surf

import (
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
)

// GetAccountAssetsParams is params to GetAccount transaction
type GetAccountAssetsParams struct {
	Address string `json:"address"`
}

// GetAccountAssetsResult is result of GetAccount
type GetAccountAssetsResult struct {
	Assets []database.Asset `json:"assets"`
}

// GetAccountAssets lookup assets of an account
func (service Service) GetAccountAssets(r *http.Request, params *GetAccountAssetsParams, result *GetAccountAssetsResult) error {
	var account database.Account
	if err := service.db.Where(database.Account{Address: params.Address}).First(&account).Error; err != nil {
		return err
	}

	var assets []database.Asset
	if err := service.db.
		Joins("Token").
		Joins("Account").
		Order("assets.id ASC").
		Where(database.Asset{AccountID: account.ID}).
		Find(&assets).Error; err != nil {
		return err
	}

	result.Assets = assets
	return nil
}
