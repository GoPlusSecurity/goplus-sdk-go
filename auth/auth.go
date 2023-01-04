package auth

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/GoPlusSecurity/goplus-sdk-go/conf"
	"net/http"
	"strconv"
	"time"
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
	now := strconv.FormatInt(time.Now().Unix(), 10)
	s := sign(a.Key, a.Secret, now)
	url := conf.Domain + "/api/v1/token"
	param := make(map[string]string)
	param["app_key"] = a.Key
	param["sign"] = s
	param["time"] = now

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

	var res Result

	err = json.NewDecoder(response.Body).Decode(&res)

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
