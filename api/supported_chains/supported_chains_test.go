package supported_chains

import (
	"testing"
)

func TestSupportedChains_Run(t *testing.T) {
	accessToken := ""
	supportedChains := NewSupportedChains(accessToken, nil)
	data, err := supportedChains.Run("")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != 1 {
		t.Errorf(data.Message)
	}
}
