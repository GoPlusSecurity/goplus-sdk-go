package chain

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type Chain struct {
	Config *Config
}

func NewChain(config *Config) *Chain {
	return &Chain{
		Config: config,
	}
}

type Result = token_controller_v_1.GetChainsListUsingGETOK

func (s *Chain) Run(name string) (*Result, error) {
	params := token_controller_v_1.NewGetChainsListUsingGETParams()
	if name != "" {
		params.SetName(pointer.String(name))
	}
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.TokenControllerv1.GetChainsListUsingGET(params)
}
