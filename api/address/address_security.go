package address

import (
	"encoding/json"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type AddressSecurity struct {
	AccessToken string
	Config      *Config
}

func NewAddressSecurity(accessToken string, config *Config) *AddressSecurity {
	return &AddressSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Result  map[string]string `json:"result"`
}

func (s *AddressSecurity) Run(chainId, address string) (*Result, error) {
	params := approve_controller_v_1.NewAddressContractUsingGET1Params()
	params.SetAddress(address)
	if chainId != "" {
		params.SetChainID(pointer.String(chainId))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.ApproveControllerv1.AddressContractUsingGET1(params)
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
