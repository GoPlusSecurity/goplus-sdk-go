package decode

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/contract_abi_controller"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type Result = contract_abi_controller.GetAbiDataInfoUsingPOSTOK

type SignatureDataDecode struct {
	Config *Config
}

func NewSignatureDataDecode(config *Config) *SignatureDataDecode {
	return &SignatureDataDecode{
		Config: config,
	}
}

func (s *SignatureDataDecode) Run(chainId, contractAddress, data string) (*Result, error) {
	params := contract_abi_controller.NewGetAbiDataInfoUsingPOSTParams()
	params.SetAbiDataRequest(&models.ParseAbiDataRequest{
		ChainID:         pointer.String(chainId),
		ContractAddress: contractAddress,
		Data:            pointer.String(data),
	})
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	response, _, err := client.Default.ContractAbiController.GetAbiDataInfoUsingPOST(params)
	return response, err
}
