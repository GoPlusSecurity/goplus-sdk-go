package token

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestTokenSecurity_Run(t *testing.T) {
	tokenSecurity := NewTokenSecurity(nil)
	data, err := tokenSecurity.Run("1", []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"})

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}

func TestSolanaTokenSecurity_Run(t *testing.T) {
	tokenSecurity := NewTokenSecurity(nil)
	data, err := tokenSecurity.SolanaTokenSecurity([]string{"HZ1JovNiVvGrGNiiYvEozEVgZ58xaU3RKwX8eACQBCt3"})

	if err != nil {
		t.Errorf(err.Error())

	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}

func TestSuiTokenSecurity_Run(t *testing.T) {
	tokenSecurity := NewTokenSecurity(nil)
	data, err := tokenSecurity.SuiTokenSecurity([]string{"0x40402a987c2f8a71b755561bfbd16c2cbb991e27e609ad148809491c32bacab9::kui::KUI"})

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
