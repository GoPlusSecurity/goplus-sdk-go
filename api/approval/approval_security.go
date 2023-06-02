package approval

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type ApprovalSecurity struct {
	AccessToken string
	Config      *Config
}

func NewApprovalSecurity(accessToken string, config *Config) *ApprovalSecurity {
	return &ApprovalSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result = approve_controller_v_1.ApprovalContractUsingGETOK

func (a *ApprovalSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewApprovalContractUsingGETParams()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if a.Config != nil && a.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(a.Config.Timeout))
	}
	if a.AccessToken != "" {
		params.SetAuthorization(pointer.String(a.AccessToken))
	}

	return client.Default.ApproveControllerv1.ApprovalContractUsingGET(params)
}
