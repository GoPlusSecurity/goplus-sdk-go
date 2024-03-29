package phishingsite

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestPhishingSiteDetection_Run(t *testing.T) {
	phishingSiteSecurity := NewPhishingSiteDetection(nil)

	url := "https://xn--cm-68s.cc/"
	data, err := phishingSiteSecurity.Run(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
