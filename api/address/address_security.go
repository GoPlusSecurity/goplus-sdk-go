package address

import (
	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_1"
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

type Result = approve_controller_v_1.AddressContractUsingGET1OK

func (s *AddressSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewAddressContractUsingGET1Params()
	params.SetAddress(address)
	if chainId != "" {
		params.SetChainID(pointer.String(chainId))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	return client.Default.ApproveControllerv1.AddressContractUsingGET1(params)
}
