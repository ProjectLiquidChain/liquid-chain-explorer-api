package liquid

import "github.com/QuoineFinancial/liquid-chain-explorer-api/common"

type API struct {
	common.RestfulAPI
}

const LiquidAPIBase = "https://api.liquid.com"

func NewAPI() API {
	return API{common.NewRestfulAPI(LiquidAPIBase)}
}
