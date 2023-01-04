package decode

import "testing"

func TestSignatureDataDecode_Run(t *testing.T) {
	accessToken := ""
	signatureDataDecode := NewSignatureDataDecode(accessToken, nil)

	chainId := "1"
	contractAddress := "0x4cc8aa0c6ffbe18534584da9b592aa438733ee66"
	inputData := "0xa0712d680000000000000000000000000000000000000000000000000000000062fee481"
	data, err := signatureDataDecode.Run(chainId, contractAddress, inputData)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != 1 {
		t.Errorf(data.Message)
	}
}
