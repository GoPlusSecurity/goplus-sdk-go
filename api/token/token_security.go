package token

import (
	"encoding/json"
	"fmt"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"strings"
	"time"
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
	url := fmt.Sprintf(conf.Domain+"/api/v1/token_security/%s?contract_addresses=%s", chainId, contractAddresses)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if s.AccessToken != "" {
		request.Header.Add("Authorization", s.AccessToken)
	}

	if s.Config != nil && s.Config.Timeout > 0 {
		client.Timeout = time.Duration(s.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	var res Result

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
