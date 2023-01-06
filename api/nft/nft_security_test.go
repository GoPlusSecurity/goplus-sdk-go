package nft

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestNFTSecurity_Run(t *testing.T) {
	accessToken := ""
	nftSecurity := NewNFTSecurity(accessToken, nil)

	chainId := "1"
	contractAddress := "0x82f5ef9ddc3d231962ba57a9c2ebb307dc8d26c2"
	data, err := nftSecurity.Run(chainId, contractAddress)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS && data.Code != errorcode.DATA_PENDING_SYNC {
		t.Errorf(data.Message)
	}

}
