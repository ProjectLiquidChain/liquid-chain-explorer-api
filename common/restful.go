package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type RestfulAPI struct {
	http.Client
	base string
}

func NewRestfulAPI(base string) RestfulAPI {
	return RestfulAPI{
		Client: http.Client{},
		base:   base,
	}
}

func (api RestfulAPI) Request(method, path string, body, result interface{}) error {
	bodyString, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, api.base+path, bytes.NewBuffer([]byte(bodyString)))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	httpResponse, err := (api.Client).Do(request)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()
	responseBody, _ := ioutil.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode > 202 {
		log.Println(method, api.base+path, string(responseBody))
		return errors.New("node error: " + string(responseBody))
	}

	return json.Unmarshal([]byte(string(responseBody)), result)
}

// Get invokes get request
func (api RestfulAPI) Get(path string, result interface{}) error {
	return api.Request("GET", path, "", result)
}

// Post invokes post request
func (api RestfulAPI) Post(path string, request interface{}, result interface{}) error {
	return api.Request("POST", path, request, result)
}
