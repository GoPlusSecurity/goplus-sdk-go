package approval

import (
	"encoding/json"
	"fmt"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"time"
)

type ConfigV2 struct {
	// timeout seconds
	Timeout int
}

type ApprovalSecurityV2 struct {
	AccessToken string
	Config      *ConfigV2
}

func NewApprovalSecurityV2(accessToken string, config *ConfigV2) *ApprovalSecurityV2 {
	return &ApprovalSecurityV2{
		AccessToken: accessToken,
		Config:      config,
	}
}

type ResultV2 struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  []map[string]any `json:"result"`
}

func (a *ApprovalSecurityV2) Token(chainId, address string) (*ResultV2, error) {
	url := fmt.Sprintf(conf.Domain+"/api/v2/token_approval_security/%s?addresses=%s", chainId, address)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if a.AccessToken != "" {
		request.Header.Add("Authorization", a.AccessToken)
	}

	if a.Config != nil && a.Config.Timeout > 0 {
		client.Timeout = time.Duration(a.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	var res ResultV2

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *ApprovalSecurityV2) ERC721NFT(chainId, address string) (*ResultV2, error) {
	url := fmt.Sprintf(conf.Domain+"/api/v2/nft721_approval_security/%s?addresses=%s", chainId, address)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if a.AccessToken != "" {
		request.Header.Add("Authorization", a.AccessToken)
	}

	if a.Config != nil && a.Config.Timeout > 0 {
		client.Timeout = time.Duration(a.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	var res ResultV2

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *ApprovalSecurityV2) ERC1155NFT(chainId, address string) (*ResultV2, error) {
	url := fmt.Sprintf(conf.Domain+"/api/v2/nft1155_approval_security/%s?addresses=%s", chainId, address)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if a.AccessToken != "" {
		request.Header.Add("Authorization", a.AccessToken)
	}

	if a.Config != nil && a.Config.Timeout > 0 {
		client.Timeout = time.Duration(a.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	var res ResultV2

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
