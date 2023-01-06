package phishing_site

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestPhishingSiteDetection_Run(t *testing.T) {
	accessToken := ""
	phishingSiteSecurity := NewPhishingSiteDetection(accessToken, nil)

	url := "https://xn--cm-68s.cc/"
	data, err := phishingSiteSecurity.Run(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}
