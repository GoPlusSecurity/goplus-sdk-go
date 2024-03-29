package rugpull

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/defi_controller"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type RugPullDetection struct {
	Config *Config
}

type Result = defi_controller.GetDefiInfoUsingGETOK

func NewRugPullDetection(config *Config) *RugPullDetection {
	return &RugPullDetection{
		Config: config,
	}
}

func (s *RugPullDetection) Run(chainId, address string) (*Result, error) {
	params := defi_controller.NewGetDefiInfoUsingGETParams()
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

	return client.Default.DefiController.GetDefiInfoUsingGET(params)
}
