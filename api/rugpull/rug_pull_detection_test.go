package rugpull

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestPhishingSiteDetection_Run(t *testing.T) {
	rugPull := NewRugPullDetection(nil)
	data, err := rugPull.Run("1", "0x6B175474E89094C44Da98b954EedeAC495271d0F")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
