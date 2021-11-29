package google

import (
	"encoding/base32"
	"github.com/dgryski/dgoogauth"
	"github.com/pkg/errors"
	"net/url"
)

const issuer = "monic-ai.com"

func CreateGoogleAuthSecret(secret string) (string, error) {
	if secret == "" {
		return "", errors.New("CreateGoogleAuthSecret account can not empty string")
	} else {
		return base32.StdEncoding.EncodeToString([]byte(secret)), nil
	}
}

func CreateGoogleAuthUrl(account string, secret string) (string, error) {
	if account == "" {
		return "", errors.New("CreateGoogleAuthSecret account can not empty string")
	} else if _url, err := url.Parse("otpauth://totp"); err != nil {
		return "", err
	} else {
		_url.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
		params := url.Values{}
		params.Add("secret", secret)
		params.Add("issuer", issuer)
		_url.RawQuery = params.Encode()
		return  _url.String(), nil
	}
}

func VerifyGoogleAuth(code string, secret string) error {
	if code == "" {
		return errors.New("VerifyGoogleAuth code can not empty string")
	} else if secret == "" {
		return errors.New("VerifyGoogleAuth secret can not empty string")
	} else {
		otpc := &dgoogauth.OTPConfig{
			Secret:      secret,
			WindowSize:  3,
			HotpCounter: 0,
		}

		val, err := otpc.Authenticate(code)
		if err != nil {
			return err
		}

		if !val {
			return errors.New("無效的Google驗證碼")
		}
		return nil
	}
}