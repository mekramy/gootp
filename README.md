# gootp

`gootp` is a Go package that provides an interface for One-Time Password (OTP) operations. It includes methods to validate an OTP code, retrieve the raw setup key, and generate a QR code image for the setup key. This package is particularly useful for implementing two-factor authentication (2FA) in your Go applications.

## Installation

To install the package, run:

```sh
go get github.com/mekramy/gootp
```

## Usage

### Creating a New Google OTP

To create a new Google OTP instance, use the `NewGoogleOTP` function. This function requires an issuer, a username, and user keys list (id, username, etc...).

```go
package main

import (
    "fmt"
    "gootp"
)

func main() {
    otp := gootp.NewGoogleOTP("MyIssuer", "MyUsername", "fasdf-32123-3123", "MyUsername")
    fmt.Println("OTP instance created")
}
```

### OTP Interface

The `OTP` interface provides the following methods:

#### Validate

Validates the provided OTP code. Returns `true` if the code is valid, otherwise `false`. An error is returned if the validation process fails.

```go
valid, err := otp.Validate("123456")
if err != nil {
    fmt.Println("Error validating OTP:", err)
} else if valid {
    fmt.Println("OTP is valid")
} else {
    fmt.Println("OTP is invalid")
}
```

#### RAW

Retrieves the raw setup key. Returns the raw setup key as a string. An error is returned if the retrieval process fails.

```go
rawKey, err := otp.RAW()
if err != nil {
    fmt.Println("Error retrieving raw key:", err)
} else {
    fmt.Println("Raw setup key:", rawKey)
}
```

#### QR

Generates a QR code image for the setup key. Returns the QR code image as a byte slice. An error is returned if the generation process fails.

```go
qrCode, err := otp.QR()
if err != nil {
    fmt.Println("Error generating QR code:", err)
} else {
    fmt.Println("QR code generated")
    // Save or display the QR code image
}
```
