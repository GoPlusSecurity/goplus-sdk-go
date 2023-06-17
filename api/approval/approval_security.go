package approval

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type ApprovalSecurity struct {
	Config *Config
}

func NewApprovalSecurity(config *Config) *ApprovalSecurity {
	return &ApprovalSecurity{
		Config: config,
	}
}

type Result = approve_controller_v_1.ApprovalContractUsingGETOK

func (a *ApprovalSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewApprovalContractUsingGETParams()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if a.Config != nil {
		if a.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(a.Config.AccessToken))
		}
		if a.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(a.Config.Timeout) * time.Second)
		}
	}

	return client.Default.ApproveControllerv1.ApprovalContractUsingGET(params)
}
