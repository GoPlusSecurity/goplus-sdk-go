package auth

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/token_controller"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"

	"k8s.io/utils/pointer"
)

type accessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Config struct {
	Timeout int
}

type Result struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  *accessTokenInfo `json:"result"`
}

type App struct {
	Key    string
	Secret string
	Config *Config
}

func NewApp(key, secret string, config *Config) *App {
	return &App{
		Key:    key,
		Secret: secret,
		Config: config,
	}
}

func (a *App) GetAccessToken() (*Result, error) {
	now := time.Now().Unix()
	nowStr := strconv.FormatInt(now, 10)
	s := sign(a.Key, a.Secret, nowStr)
	params := token_controller.NewGetAccessTokenUsingPOSTParams()
	params.SetRequest(&models.GetAccessTokenRequest{
		AppKey: pointer.String(a.Key),
		Sign:   pointer.String(s),
		Time:   pointer.Int64(now),
	})
	if a.Config != nil && a.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(a.Config.Timeout))
	}

	response, _, err := client.Default.TokenController.GetAccessTokenUsingPOST(params)
	if err != nil {
		return nil, err
	}

	tmp, err := json.Marshal(response.Payload)
	if err != nil {
		return nil, err
	}

	res := Result{}
	err = json.Unmarshal(tmp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func sign(key, secret, time string) string {
	h := sha1.New()
	s := key + time + secret
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
