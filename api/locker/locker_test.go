package locker

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestLocker_LockerToken(t *testing.T) {
	locker := NewLocker(nil)
	data, err := locker.LockerToken("8453", "0x6fd0303649296360f10c07b24521deda9223086d", 1, 100)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}

func TestLocker_LockerLPV3(t *testing.T) {
	locker := NewLocker(nil)
	data, err := locker.LockerLPV3("56", "0x579df956c6cE6178fBBD78bbE4f05786cFBA9B76", 1, 100)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS {
		t.Errorf(data.Payload.Message)
	}
}
