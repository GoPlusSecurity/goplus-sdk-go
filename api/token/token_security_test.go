package token

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestTokenSecurity_Run(t *testing.T) {
	tokenSecurity := NewTokenSecurity("", nil)
	data, err := tokenSecurity.Run("1", []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"})

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}
