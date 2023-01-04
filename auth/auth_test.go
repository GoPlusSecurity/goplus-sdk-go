package auth

import "testing"

func TestApp_GetAccessToken(t *testing.T) {
	key := ""
	secret := ""
	app := NewApp(key, secret, nil)
	_, err := app.GetAccessToken()

	if err != nil {
		t.Errorf(err.Error())
	}
}
