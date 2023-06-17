package nft

import (
	"time"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/nft_controller"
	"k8s.io/utils/pointer"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type NFTSecurity struct {
	Config *Config
}

func NewNFTSecurity(config *Config) *NFTSecurity {
	return &NFTSecurity{
		Config: config,
	}
}

type Result = nft_controller.GetNftInfoUsingGET1OK

func (s *NFTSecurity) Run(chainId, address string) (*Result, error) {
	params := nft_controller.NewGetNftInfoUsingGET1Params()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.NftController.GetNftInfoUsingGET1(params)
}
