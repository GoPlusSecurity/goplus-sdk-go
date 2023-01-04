package phishing_site

import (
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

	if data.Code != 1 {
		t.Errorf(data.Message)
	}
}
