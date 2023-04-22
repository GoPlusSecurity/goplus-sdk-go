package phishing_site

import (
	"encoding/json"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/website_controller"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type PhishingSiteDetection struct {
	AccessToken string
	Config      *Config
}

func NewPhishingSiteDetection(accessToken string, config *Config) *PhishingSiteDetection {
	return &PhishingSiteDetection{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (s *PhishingSiteDetection) Run(url string) (*Result, error) {
	params := website_controller.NewPhishingSiteUsingGETParams()
	params.SetURL(url)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.WebsiteController.PhishingSiteUsingGET(params)
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
