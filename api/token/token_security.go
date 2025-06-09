package token

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_security_api_for_solana_beta"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_security_api_for_sui"
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

type SolanaTokenSecurityResult = token_security_api_for_solana_beta.SolanaTokenSecurityUsingGETOK

func (s *TokenSecurity) SolanaTokenSecurity(addresses []string) (*SolanaTokenSecurityResult, error) {
	contractAddresses := strings.Join(addresses, ",")

	params := token_security_api_for_solana_beta.NewSolanaTokenSecurityUsingGETParams()
	params.SetContractAddresses(contractAddresses)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.TokenSecurityAPIForSolanaBeta.SolanaTokenSecurityUsingGET(params)
}

type SuiTokenSecurityResult = token_security_api_for_sui.SuiTokenSecurityUsingGETOK

func (s *TokenSecurity) SuiTokenSecurity(addresses []string) (*SuiTokenSecurityResult, error) {
	contractAddresses := strings.Join(addresses, ",")

	params := token_security_api_for_sui.NewSuiTokenSecurityUsingGETParams()
	params.SetContractAddresses(contractAddresses)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.TokenSecurityAPIForSui.SuiTokenSecurityUsingGET(params)
}
