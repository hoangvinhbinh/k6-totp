package totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"go.k6.io/k6/js/modules"
)

// init registers the TOTP module with k6
func init() {
	modules.Register("k6/x/totp", new(TOTP))
}

// TOTP structure for generating TOTP tokens
type TOTP struct{}

// Generate function to create a TOTP token based on a secret
func (t *TOTP) Generate(secret string) (string, error) {
	fmt.Println("Generate method called") // Debugging line
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
