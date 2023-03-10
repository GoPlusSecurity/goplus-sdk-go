package chain

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestChain_Run(t *testing.T) {
	accessToken := ""
	chain := NewChain(accessToken, nil)
	data, err := chain.Run("")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}
