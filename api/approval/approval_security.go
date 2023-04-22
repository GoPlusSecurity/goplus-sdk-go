package approval

import (
	"encoding/json"
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

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (a *ApprovalSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewApprovalContractUsingGETParams()
	params.SetChainID(chainId)
	params.SetContractAddresses(pointer.String(address))
	if a.Config != nil && a.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(a.Config.Timeout))
	}
	if a.AccessToken != "" {
		params.SetAuthorization(pointer.String(a.AccessToken))
	}

	response, err := client.Default.ApproveControllerv1.ApprovalContractUsingGET(params)
	if err != nil {
		return nil, err
	}

	tmp, err := json.Marshal(response.Payload)
	if err != nil {
		return nil, err
	}

	res := Result{}
	err = json.Unmarshal(tmp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
