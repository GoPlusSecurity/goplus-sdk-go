package dapp

import (
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

type Result = dapp_controller.GetDappInfoUsingGETOK

func (s *DAppSecurity) Run(url string) (*Result, error) {
	params := dapp_controller.NewGetDappInfoUsingGETParams()
	params.SetURL(pointer.String(url))
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	return client.Default.DappController.GetDappInfoUsingGET(params)
}
