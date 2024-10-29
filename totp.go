package totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"go.k6.io/k6/js/modules"
)

// TOTP structure
type TOTP struct{}

// init function to register the TOTP module
func init() {
	modules.Register("k6/x/totp", new(TOTP))
}

// Generate function to create a TOTP token based on a secret
func (t *TOTP) Generate(secret string) (string, error) {
	fmt.Println("Generate method called") // Debugging line
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
