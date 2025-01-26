package gootp

import (
	"net/url"

	"github.com/dgryski/dgoogauth"
	"rsc.io/qr"
)

type googleOTP struct {
	issuer     string
	username   string
	secret     string
	googleAuth *dgoogauth.OTPConfig
}

func (driver *googleOTP) Validate(code string) (bool, error) {
	return driver.googleAuth.Authenticate(code)
}

func (driver *googleOTP) RAW() (string, error) {
	URL, err := url.Parse("otpauth://totp/" + url.PathEscape(driver.username))
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("secret", driver.secret)
	params.Add("issuer", driver.issuer)
	params.Add("digits", "6")

	URL.RawQuery = params.Encode()
	return URL.String(), nil
}

func (driver *googleOTP) QR() ([]byte, error) {
	raw, err := driver.RAW()
	if err != nil {
		return nil, err
	}

	code, err := qr.Encode(raw, qr.Q)
	if err != nil {
		return nil, err
	}

	return code.PNG(), nil
}
