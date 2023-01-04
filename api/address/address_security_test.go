package address

import "testing"

func TestAddressSecurity_Run(t *testing.T) {
	accessToken := ""
	addressSecurity := NewAddressSecurity(accessToken, nil)
	data, err := addressSecurity.Run("", "0xc8b759860149542a98a3eb57c14aadf59d6d89b9")
	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Code != 1 {
		t.Errorf(data.Message)
	}
}
