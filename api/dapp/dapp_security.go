package dapp

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/dapp_controller"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type DAppSecurity struct {
	Config *Config
}

func NewDAppSecurity(config *Config) *DAppSecurity {
	return &DAppSecurity{
		Config: config,
	}
}

type Result = dapp_controller.GetDappInfoUsingGETOK

func (s *DAppSecurity) Run(url string) (*Result, error) {
	params := dapp_controller.NewGetDappInfoUsingGETParams()
	params.SetURL(pointer.String(url))
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.DappController.GetDappInfoUsingGET(params)
}
