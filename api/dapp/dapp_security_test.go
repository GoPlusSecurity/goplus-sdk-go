package dapp

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestDAppSecurity_Run(t *testing.T) {
	accessToken := ""
	dAppSecurity := NewDAppSecurity(accessToken, nil)

	url := "https://for.tube"
	data, err := dAppSecurity.Run(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}

}
