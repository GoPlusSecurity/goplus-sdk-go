package phishing_site

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/website_controller"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type PhishingSiteDetection struct {
	Config *Config
}

func NewPhishingSiteDetection(config *Config) *PhishingSiteDetection {
	return &PhishingSiteDetection{
		Config: config,
	}
}

type Result = website_controller.PhishingSiteUsingGETOK

func (s *PhishingSiteDetection) Run(url string) (*Result, error) {
	params := website_controller.NewPhishingSiteUsingGETParams()
	params.SetURL(url)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}

	return client.Default.WebsiteController.PhishingSiteUsingGET(params)
}
