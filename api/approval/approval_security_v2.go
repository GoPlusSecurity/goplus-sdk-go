package approval

import (
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/approve_controller_v_2"
)

type ConfigV2 struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type ApprovalSecurityV2 struct {
	Config *ConfigV2
}

func NewApprovalSecurityV2(config *ConfigV2) *ApprovalSecurityV2 {
	return &ApprovalSecurityV2{
		Config: config,
	}
}

type TokenResult = approve_controller_v_2.AddressTokenApproveListUsingGET1OK

func (a *ApprovalSecurityV2) Token(chainId, address string) (*TokenResult, error) {
	params := approve_controller_v_2.NewAddressTokenApproveListUsingGET1Params()
	params.SetChainID(chainId)
	params.SetAddresses(address)
	if a.Config != nil {
		if a.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(a.Config.AccessToken))
		}
		if a.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(a.Config.Timeout) * time.Second)
		}
	}

	return client.Default.ApproveControllerv2.AddressTokenApproveListUsingGET1(params)
}

type ERC721NFTResult = approve_controller_v_2.AddressNFT721ApproveListUsingGET1OK

func (a *ApprovalSecurityV2) ERC721NFT(chainId, address string) (*ERC721NFTResult, error) {
	params := approve_controller_v_2.NewAddressNFT721ApproveListUsingGET1Params()
	params.SetChainID(chainId)
	params.SetAddresses(address)
	if a.Config != nil {
		if a.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(a.Config.AccessToken))
		}
		if a.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(a.Config.Timeout) * time.Second)
		}
	}

	return client.Default.ApproveControllerv2.AddressNFT721ApproveListUsingGET1(params)
}

type ERC1155NFTResult = approve_controller_v_2.AddressNFT1155ApproveListUsingGET1OK

func (a *ApprovalSecurityV2) ERC1155NFT(chainId, address string) (*ERC1155NFTResult, error) {
	params := approve_controller_v_2.NewAddressNFT1155ApproveListUsingGET1Params()
	params.SetChainID(chainId)
	params.SetAddresses(address)
	if a.Config != nil {
		if a.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(a.Config.AccessToken))
		}
		if a.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(a.Config.Timeout) * time.Second)
		}
	}

	return client.Default.ApproveControllerv2.AddressNFT1155ApproveListUsingGET1(params)
}
