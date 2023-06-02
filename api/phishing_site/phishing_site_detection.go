package phishing_site

import (
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

type Result = website_controller.PhishingSiteUsingGETOK

func (s *PhishingSiteDetection) Run(url string) (*Result, error) {
	params := website_controller.NewPhishingSiteUsingGETParams()
	params.SetURL(url)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	return client.Default.WebsiteController.PhishingSiteUsingGET(params)
}
