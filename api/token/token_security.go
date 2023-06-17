package token

import (
	"strings"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type TokenSecurity struct {
	Config *Config
}

func NewTokenSecurity(config *Config) *TokenSecurity {
	return &TokenSecurity{
		Config: config,
	}
}

type Result = token_controller_v_1.TokenSecurityUsingGET1OK

func (s *TokenSecurity) Run(chainId string, addresses []string) (*Result, error) {
	contractAddresses := strings.Join(addresses, ",")

	params := token_controller_v_1.NewTokenSecurityUsingGET1Params()
	params.SetChainID(chainId)
	params.SetContractAddresses(contractAddresses)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.TokenControllerv1.TokenSecurityUsingGET1(params)
}
