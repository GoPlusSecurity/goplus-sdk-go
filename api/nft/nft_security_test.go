package nft

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestNFTSecurity_Run(t *testing.T) {
	nftSecurity := NewNFTSecurity(nil)

	chainId := "1"
	contractAddress := "0x82f5ef9ddc3d231962ba57a9c2ebb307dc8d26c2"
	data, err := nftSecurity.Run(chainId, contractAddress)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS && data.Payload.Code != errorcode.DATA_PENDING_SYNC {
		t.Errorf(data.Payload.Message)
	}

}
