package xk6_totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"go.k6.io/k6/js/modules"
)

// init function to register the TOTP module
func init() {
	fmt.Println("Initializing xk6 TOTP module")
	modules.Register("k6/x/xk6_totp", new(TOTP))
}

// TOTP structure
type TOTP struct{}

// Generate generates a TOTP token for the given secret.
func (t *TOTP) Generate(secret string) (string, error) {
	fmt.Println("Generate method called") // Debugging line
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
