package tx_simulation

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/transaction_simulation_for_solana"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type TxSimulation struct {
	Config *Config
}

func NewTxSimulation(config *Config) *TxSimulation {
	return &TxSimulation{
		Config: config,
	}
}

type SolanaTxSimulationResult = transaction_simulation_for_solana.PrerunTxUsingPOSTOK

func (s *TxSimulation) SolanaTxSimulation(encodedTransaction string) (*SolanaTxSimulationResult, error) {
	params := transaction_simulation_for_solana.NewPrerunTxUsingPOSTParams()
	params.SetRequest(&models.SolanaPrerunTxRequest{EncodedTransaction: encodedTransaction})
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	response, _, err := client.Default.TransactionSimulationForSolana.PrerunTxUsingPOST(params)

	return response, err
}
