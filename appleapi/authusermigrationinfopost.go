package appleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	AUTH_USERMIGRATIONINFO_URL = `https://appleid.apple.com/auth/usermigrationinfo`
)

type UserMigrationInfoForTransferSubResp struct {
	TransferSub string `json:"transfer_sub"`
}

func AuthUsermigrationinfoPOSTForTransferSub(accessToken, sub, target, clientID, clientSecret string) (*UserMigrationInfoForTransferSubResp, error) {
	req, err := http.NewRequest(http.MethodPost, AUTH_USERMIGRATIONINFO_URL, strings.NewReader(url.Values{
		"sub":           {sub},
		"target":        {target},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respStruct := &UserMigrationInfoForTransferSubResp{}

	if err := json.Unmarshal(b, &respStruct); err != nil {
		return nil, err
	}

	if respStruct.TransferSub == "" {
		return nil, fmt.Errorf(`transfer_sub == "", resp: %s`, string(b))
	}

	return respStruct, nil
}

type UserMigrationInfoForSubResp struct {
	Sub            string `json:"sub"`
	Email          string `json:"email"`
	IsPrivateEmail bool   `json:"is_private_email" `
}

func AuthUsermigrationinfoPOSTForSub(accessToken, transferSub, clientID, clientSecret string) (*UserMigrationInfoForSubResp, error) {
	req, err := http.NewRequest(http.MethodPost, AUTH_USERMIGRATIONINFO_URL, strings.NewReader(url.Values{
		"transfer_sub":  {transferSub},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respStruct := &UserMigrationInfoForSubResp{}

	if err := json.Unmarshal(b, &respStruct); err != nil {
		return nil, err
	}

	if respStruct.Sub == "" {
		return nil, fmt.Errorf(`sub == "", resp: %s`, string(b))
	}

	return respStruct, nil
}
