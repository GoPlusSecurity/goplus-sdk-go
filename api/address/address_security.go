package address

import (
	"encoding/json"
	"fmt"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"time"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type AddressSecurity struct {
	AccessToken string
	Config      *Config
}

func NewAddressSecurity(accessToken string, config *Config) *AddressSecurity {
	return &AddressSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Result  map[string]string `json:"result"`
}

func (s *AddressSecurity) Run(chainId, address string) (*Result, error) {
	url := fmt.Sprintf(conf.Domain+"/api/v1/address_security/%s", address)

	if chainId != "" {
		url = url + "?chain_id" + chainId
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if s.AccessToken != "" {
		request.Header.Add("Authorization", s.AccessToken)
	}

	if s.Config != nil && s.Config.Timeout > 0 {
		client.Timeout = time.Duration(s.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	var res Result

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
