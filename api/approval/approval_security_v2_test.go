package approval

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"testing"
)

func TestApprovalSecurityV2_Token(t *testing.T) {
	accessToken := ""
	approvalSecurityV2 := NewApprovalSecurityV2(accessToken, nil)
	data, err := approvalSecurityV2.Token("56", "0xd018e2b543a2669410537f96293590138cacedf3")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}

func TestApprovalSecurityV2_ERC721NFT(t *testing.T) {

	accessToken := ""
	approvalSecurityV2 := NewApprovalSecurityV2(accessToken, nil)
	data, err := approvalSecurityV2.ERC721NFT("1", "0xd95dbdab08a9fed2d71ac9c3028aac40905d8cf3")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}

func TestApprovalSecurityV2_ERC1155NFT(t *testing.T) {

	accessToken := ""
	approvalSecurityV2 := NewApprovalSecurityV2(accessToken, nil)
	data, err := approvalSecurityV2.ERC1155NFT("56", "0xb0dccbb9c4a65a94a41a0165aaea79c8b2fc54ce")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != errorcode.SUCCESS {
		t.Errorf(data.Message)
	}
}
