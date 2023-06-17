package approval

import (
	"testing"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
)

func TestApprovalSecurity_Run(t *testing.T) {
	approvalSecurityV1 := NewApprovalSecurity(nil)
	data, err := approvalSecurityV1.Run("1", "0x4639cd8cd52ec1cf2e496a606ce28d8afb1c792f")

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Payload.Code != errorcode.SUCCESS && data.Payload.Code != errorcode.DATA_PENDING_SYNC {
		t.Errorf(data.Payload.Message)
	}
}
