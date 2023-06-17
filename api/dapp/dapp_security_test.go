package dapp

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestDAppSecurity_Run(t *testing.T) {
	dAppSecurity := NewDAppSecurity(nil)

	url := "https://for.tube"
	data, err := dAppSecurity.Run(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}

}
