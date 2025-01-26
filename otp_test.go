package gootp_test

import (
	"testing"

	"github.com/mekramy/gootp"
)

func TestOTP(t *testing.T) {

	issuer := "testIssuer"
	username := "testUser"
	keys := []string{"key1", "key2"}

	otp := gootp.NewGoogleOTP(issuer, username, keys...)

	// Test RAW method
	raw, err := otp.RAW()
	if err != nil {
		t.Fatalf("Failed to get RAW key: %v", err)
	}
	if raw == "" {
		t.Error("Expected non-empty RAW key")
	}

	// Test Validate method with an invalid code
	valid, err := otp.Validate("123456")
	if err != nil {
		t.Fatalf("Failed to validate OTP: %v", err)
	}
	if valid {
		t.Error("Expected invalid OTP code")
	}

	// Test QR method
	qr, err := otp.QR()
	if err != nil {
		t.Fatalf("Failed to generate QR code: %v", err)
	}
	if len(qr) == 0 {
		t.Error("Expected non-empty QR code")
	}
}
