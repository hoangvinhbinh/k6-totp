package totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"go.k6.io/k6/js/modules"
)

// init function to register the TOTP module
func init() {
	fmt.Println("Initializing TOTP module")
	modules.Register("k6/x/totp", new(TOTP))
}

// TOTP structure
type TOTP struct{}

func (*TOTP) Hello(name string) string {
	return "Hello " + name
}

// Generate function to create a TOTP token based on a secret
func (*TOTP) Generate(secret string) (string, error) {
	fmt.Println("Generate method called") // Debugging line
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
