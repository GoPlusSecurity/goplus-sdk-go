package locker

import (
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/lock_controller"
	"time"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"k8s.io/utils/pointer"
)

type Config struct {
	// timeout seconds
	Timeout     int
	AccessToken string
}

type Locker struct {
	Config *Config
}

func NewLocker(config *Config) *Locker {
	return &Locker{
		Config: config,
	}
}

type LockerTokenResult = lock_controller.GetTokenLockersUsingGETOK

func (s *Locker) LockerToken(chainId string, tokenAddress string, pageNum int32, pageSize int32) (*LockerTokenResult, error) {

	params := lock_controller.NewGetTokenLockersUsingGETParams()
	params.SetChainID(chainId)
	params.SetTokenAddress(tokenAddress)
	params.SetPageNum(pageNum)
	params.SetPageSize(pageSize)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}
	return client.Default.LockController.GetTokenLockersUsingGET(params)
}

type LockerLPV3Result = lock_controller.GetNftLockersUsingGETOK

func (s *Locker) LockerLPV3(chainId string, poolAddress string, pageNum int32, pageSize int32) (*LockerLPV3Result, error) {

	params := lock_controller.NewGetNftLockersUsingGETParams()
	params.SetChainID(chainId)
	params.SetPoolAddress(poolAddress)
	params.SetPageNum(pageNum)
	params.SetPageSize(pageSize)
	if s.Config != nil {
		if s.Config.AccessToken != "" {
			params.SetAuthorization(pointer.String(s.Config.AccessToken))
		}
		if s.Config.Timeout != 0 {
			params.SetTimeout(time.Duration(s.Config.Timeout) * time.Second)
		}
	}
	return client.Default.LockController.GetNftLockersUsingGET(params)
}
