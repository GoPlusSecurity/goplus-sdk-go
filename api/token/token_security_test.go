package token

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestTokenSecurity_Run(t *testing.T) {
	tokenSecurity := NewTokenSecurity("", nil)
	data, err := tokenSecurity.Run("1", []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"})

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
