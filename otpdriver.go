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

func (g *googleOTP) Validate(code string) (bool, error) {
	return g.googleAuth.Authenticate(code)
}

func (g *googleOTP) RAW() (string, error) {
	URL, err := url.Parse("otpauth://totp/" + url.PathEscape(g.username))
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("secret", g.secret)
	params.Add("issuer", g.issuer)
	params.Add("digits", "6")

	URL.RawQuery = params.Encode()
	return URL.String(), nil
}

func (g *googleOTP) QR() ([]byte, error) {
	raw, err := g.RAW()
	if err != nil {
		return nil, err
	}

	code, err := qr.Encode(raw, qr.Q)
	if err != nil {
		return nil, err
	}

	return code.PNG(), nil
}
