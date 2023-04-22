package dapp

import (
	"encoding/json"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/dapp_controller"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type DAppSecurity struct {
	AccessToken string
	Config      *Config
}

func NewDAppSecurity(accessToken string, config *Config) *DAppSecurity {
	return &DAppSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (s *DAppSecurity) Run(url string) (*Result, error) {
	params := dapp_controller.NewGetDappInfoUsingGETParams()
	params.SetURL(pointer.String(url))
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.DappController.GetDappInfoUsingGET(params)
	if err != nil {
		return nil, err
	}

	tmp, err := json.Marshal(response.Payload)
	if err != nil {
		return nil, err
	}

	res := Result{}
	err = json.Unmarshal(tmp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
