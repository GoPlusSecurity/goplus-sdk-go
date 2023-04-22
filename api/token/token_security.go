package token

import (
	"encoding/json"
	"strings"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type TokenSecurity struct {
	AccessToken string
	Config      *Config
}

func NewTokenSecurity(accessToken string, config *Config) *TokenSecurity {
	return &TokenSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Result  map[string]map[string]any `json:"result"`
}

func (s *TokenSecurity) Run(chainId string, addresses []string) (*Result, error) {
	contractAddresses := strings.Join(addresses, ",")

	params := token_controller_v_1.NewTokenSecurityUsingGET1Params()
	params.SetChainID(chainId)
	params.SetContractAddresses(contractAddresses)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.TokenControllerv1.TokenSecurityUsingGET1(params)
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
