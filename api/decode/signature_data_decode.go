package decode

import (
	"bytes"
	"encoding/json"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"time"
)

type Config struct {
	Timeout int
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

type SignatureDataDecode struct {
	AccessToken string
	Config      *Config
}

func NewSignatureDataDecode(accessToken string, config *Config) *SignatureDataDecode {
	return &SignatureDataDecode{
		AccessToken: accessToken,
		Config:      config,
	}
}

func (s *SignatureDataDecode) Run(chainId, contractAddress, data string) (*Result, error) {

	url := conf.Domain + "/api/v1/abi/input_decode"
	param := map[string]string{
		"chain_id":         chainId,
		"contract_address": contractAddress,
		"data":             data,
	}

	bytesData, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	if s.Config != nil && s.Config.Timeout > 0 {
		client.Timeout = time.Duration(s.Config.Timeout) * time.Second
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

	var res Result

	err = json.NewDecoder(response.Body).Decode(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
