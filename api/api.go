package api

import (
	"github.com/QuoineFinancial/liquid-chain-explorer-api/common"
)

type NodeAPI struct {
	common.RPCAPI
}

func New(url string) NodeAPI {
	return NodeAPI{common.NewRPCAPI(url)}
}

func (api NodeAPI) GetBlock(height uint64) (block, error) {
	var response struct {
		Block block `json:"block"`
	}
	if err := api.Request("chain.GetBlock", struct {
		Height uint64 `json:"height"`
	}{height}, &response); err != nil {
		return block{}, err
	}
	return response.Block, nil
}
