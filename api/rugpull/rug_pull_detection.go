package phishing_site

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/defi_controller"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type RugPullDetection struct {
	AccessToken string
	Config      *Config
}

type Result = defi_controller.GetDefiInfoUsingGETOK

func NewRugPullDetection(accessToken string, config *Config) *RugPullDetection {
	return &RugPullDetection{
		AccessToken: accessToken,
		Config:      config,
	}
}

func (s *RugPullDetection) Run(chainId, address string) (*Result, error) {
	params := defi_controller.NewGetDefiInfoUsingGETParams()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	return client.Default.DefiController.GetDefiInfoUsingGET(params)
}
