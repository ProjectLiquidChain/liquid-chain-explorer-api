package common

type RPCAPI struct {
	rest RestfulAPI
}

func NewRPCAPI(base string) RPCAPI {
	return RPCAPI{NewRestfulAPI(base)}
}

func (api RPCAPI) Request(method string, params interface{}, response interface{}) error {
	request := struct {
		ID      int         `json:"id"`
		Method  string      `json:"method"`
		JSONRPC string      `json:"jsonrpc"`
		Params  interface{} `json:"params"`
	}{
		ID:      1,
		Method:  method,
		JSONRPC: "2.0",
		Params:  params,
	}

	if err := api.rest.Post("", request, &response); err != nil {
		return err
	}
	return nil
}
