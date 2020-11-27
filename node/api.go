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
		Result struct {
			Block Block `json:"block"`
		} `json:"result"`
	}
	if err := api.Request("chain.GetBlockByHeight", struct {
		Height uint64 `json:"height"`
	}{height}, &response); err != nil {
		return Block{}, err
	}
	return response.Result.Block, nil
}

// GetLatestBlock returns block with given height
func (api API) GetLatestBlock() (Block, error) {
	var response struct {
		Result struct {
			Block Block `json:"block"`
		} `json:"result"`
	}
	if err := api.Request("chain.GetLatestBlock", struct{}{}, &response); err != nil {
		return Block{}, err
	}
	return response.Result.Block, nil
}

// Call emit call
func (api API) Call(method, address string, args []string) (Receipt, error) {
	var response struct {
		Result Receipt `json:"result"`
	}
	if err := api.Request("chain.Call", struct {
		Method  string   `json:"method"`
		Address string   `json:"address"`
		Args    []string `json:"args"`
	}{method, address, args}, &response); err != nil {
		return Receipt{}, err
	}
	return response.Result, nil
}
