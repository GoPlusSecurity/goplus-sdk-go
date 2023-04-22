package chain

import (
	"encoding/json"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_controller_v_1"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type Chain struct {
	AccessToken string
	Config      *Config
}

func NewChain(accessToken string, config *Config) *Chain {
	return &Chain{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Result  []map[string]string `json:"result"`
}

func (s *Chain) Run(name string) (*Result, error) {
	params := token_controller_v_1.NewGetChainsListUsingGETParams()
	if name != "" {
		params.SetName(pointer.String(name))
	}
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.TokenControllerv1.GetChainsListUsingGET(params)
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
