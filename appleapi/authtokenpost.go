package appleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	AUTH_TOKEN_URL = `https://appleid.apple.com/auth/token`
)

type AuthTokenPOSTResp struct {
	AccessToken string `json:"access_token"`
}

func AuthTokenPOST(clientID, clientSecret string) (*AuthTokenPOSTResp, error) {
	resp, err := http.PostForm(AUTH_TOKEN_URL, url.Values{
		"grant_type":    {"client_credentials"},
		"scope":         {"user.migration"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respStruct := &AuthTokenPOSTResp{}

	if err := json.Unmarshal(b, respStruct); err != nil {
		return nil, err
	}

	if respStruct.AccessToken == "" {
		return nil, fmt.Errorf(`access_token == "", resp: %s`, string(b))
	}

	return respStruct, nil
}
