package gootp

import (
	"crypto/md5"
	"encoding/base32"
	"strings"

	"github.com/dgryski/dgoogauth"
)

// OTP represents an interface for One-Time Password (OTP) operations.
// It provides methods to validate an OTP code, retrieve the raw setup key,
// and generate a QR code image for the setup key.
type OTP interface {
	// Validate checks if the provided OTP code is valid.
	// Returns true if the code is valid, otherwise false.
	// An error is returned if the validation process fails.
	Validate(code string) (bool, error)

	// RAW retrieves the raw setup key.
	// Returns the raw setup key as a string.
	// An error is returned if the retrieval process fails.
	RAW() (string, error)

	// QR generates a QR code image for the setup key.
	// Returns the QR code image as a byte slice.
	// An error is returned if the generation process fails.
	QR() ([]byte, error)
}

// NewGoogleOTP create new GoogleAuthenticator otp driver.
func NewGoogleOTP(issuer, username string, keys ...string) OTP {
	otp := new(googleOTP)
	otp.issuer = issuer
	otp.username = username

	// Generate Secret
	secret := md5.Sum([]byte(issuer + username + strings.Join(keys, "")))
	_len := len(secret)
	_mid := _len / 2
	otp.secret = base32.StdEncoding.EncodeToString([]byte{secret[0], secret[1], secret[2], secret[_mid-1], secret[_mid], secret[_mid+1], secret[_mid+2], secret[_len-3], secret[_len-2], secret[_len-1]})

	// generate driver
	otp.googleAuth = &dgoogauth.OTPConfig{
		Secret:      otp.secret,
		WindowSize:  3,
		HotpCounter: 0,
		UTC:         true,
	}
	return otp
}
