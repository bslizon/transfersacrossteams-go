package converter

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/bslizon/transfersacrossteams-go/appleapi"
	"github.com/golang-jwt/jwt"
)

type Client struct {
	TeamID   string
	KeyID    string
	KeyPEM   string
	ClientID string

	key          any
	clientSecret string
	accessToken  string
}

func New(teamID string, keyID string, KeyPEM string, clientID string) *Client {
	return &Client{
		TeamID:   teamID,
		KeyID:    keyID,
		KeyPEM:   KeyPEM,
		ClientID: clientID,
	}
}

func (c *Client) Init() error {
	return c.Refresh()
}

func (c *Client) Refresh() error {
	p, _ := pem.Decode([]byte(c.KeyPEM))
	if p == nil {
		return fmt.Errorf("pem.Decode(KeyPEM) failed")
	}

	key, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": c.TeamID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"aud": "https://appleid.apple.com",
		"sub": c.ClientID,
	})

	token.Header["kid"] = c.KeyID

	clientSecret, err := token.SignedString(key)
	if err != nil {
		return err
	}

	resp, err := appleapi.AuthTokenPOST(c.ClientID, clientSecret)
	if err != nil {
		return err
	}

	c.key = key
	c.clientSecret = clientSecret
	c.accessToken = resp.AccessToken

	return nil
}

func (c *Client) FromSubToTransferSub(sub string, target string) (string, error) {
	resp, err := appleapi.AuthUsermigrationinfoPOSTForTransferSub(c.accessToken, sub, target, c.ClientID, c.clientSecret)
	if err != nil {
		return "", err
	}

	return resp.TransferSub, nil
}

func (c *Client) FromTransferSubToSub(transferSub string) (string, error) {
	resp, err := appleapi.AuthUsermigrationinfoPOSTForSub(c.accessToken, transferSub, c.ClientID, c.clientSecret)
	if err != nil {
		return "", err
	}

	return resp.Sub, nil
}
