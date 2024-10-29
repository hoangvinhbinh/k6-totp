package totp

import (
	"time"

	"github.com/pquerna/otp/totp"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/totp", new(TOTP))
}

type TOTP struct{}

func (t *TOTP) Generate(secret string) (string, error) {
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
