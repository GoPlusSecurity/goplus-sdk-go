package chain

import (
	"testing"
)

func TestChain_Run(t *testing.T) {
	accessToken := ""
	chain := NewChain(accessToken, nil)
	data, err := chain.Run("")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != 1 {
		t.Errorf(data.Message)
	}
}
