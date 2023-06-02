package address

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestAddressSecurity_Run(t *testing.T) {
	accessToken := ""
	addressSecurity := NewAddressSecurity(accessToken, nil)
	data, err := addressSecurity.Run("", "0xc8b759860149542a98a3eb57c14aadf59d6d89b9")
	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
