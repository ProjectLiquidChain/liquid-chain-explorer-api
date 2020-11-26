package node

import (
	"github.com/QuoineFinancial/liquid-chain-explorer-api/common"
)

// API wraps methods to fetch data from node
type API struct {
	common.RPCAPI
}

// New returns new instance of NodeAPI with given url
func New(url string) API {
	return API{common.NewRPCAPI(url)}
}

// GetBlock returns block with given height
func (api API) GetBlock(height uint64) (Block, error) {
	var response struct {
		Block Block `json:"block"`
	}
	if err := api.Request("chain.GetBlock", struct {
		Height uint64 `json:"height"`
	}{height}, &response); err != nil {
		return Block{}, err
	}
	return response.Block, nil
}
