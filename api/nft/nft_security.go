package nft

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/nft_controller"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type NFTSecurity struct {
	AccessToken string
	Config      *Config
}

func NewNFTSecurity(accessToken string, config *Config) *NFTSecurity {
	return &NFTSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result = nft_controller.GetNftInfoUsingGET1OK

func (s *NFTSecurity) Run(chainId, address string) (*Result, error) {
	params := nft_controller.NewGetNftInfoUsingGET1Params()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	return client.Default.NftController.GetNftInfoUsingGET1(params)
}
