package supported_chains

import (
	"encoding/json"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"time"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type SupportedChains struct {
	AccessToken string
	Config      *Config
}

func NewSupportedChains(accessToken string, config *Config) *SupportedChains {
	return &SupportedChains{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Result  []map[string]string `json:"result"`
}

func (s *SupportedChains) Run(name string) (*Result, error) {

	url := conf.Domain + "/api/v1/supported_chains"

	if name != "" {
		url = url + "?name" + name
	}

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
