package decode

import (
	"encoding/json"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/contract_abi_controller"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

type Config struct {
	Timeout int
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

type SignatureDataDecode struct {
	AccessToken string
	Config      *Config
}

func NewSignatureDataDecode(accessToken string, config *Config) *SignatureDataDecode {
	return &SignatureDataDecode{
		AccessToken: accessToken,
		Config:      config,
	}
}

func (s *SignatureDataDecode) Run(chainId, contractAddress, data string) (*Result, error) {
	params := contract_abi_controller.NewGetAbiDataInfoUsingPOSTParams()
	params.SetAbiDataRequest(&models.ParseAbiDataRequest{
		ChainID:         pointer.String(chainId),
		ContractAddress: contractAddress,
		Data:            pointer.String(data),
	})
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, _, err := client.Default.ContractAbiController.GetAbiDataInfoUsingPOST(params)
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
