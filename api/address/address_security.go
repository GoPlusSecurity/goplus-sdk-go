package address

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout     int64
	AccessToken string
}

type AddressSecurity struct {
	Config *Config
}

func NewAddressSecurity(config *Config) *AddressSecurity {
	return &AddressSecurity{
		Config: config,
	}
}

type Result = approve_controller_v_1.AddressContractUsingGET1OK

func (s *AddressSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewAddressContractUsingGET1Params()
	params.SetAddress(address)
	if chainId != "" {
		params.SetChainID(pointer.String(chainId))
	}
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.ApproveControllerv1.AddressContractUsingGET1(params)
}
