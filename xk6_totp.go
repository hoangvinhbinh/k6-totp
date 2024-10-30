package xk6_totp

import (
	"fmt"

	"go.k6.io/k6/js/modules"
)

// init function to register the TOTP module
func init() {
	fmt.Println("Initializing xk6 TOTP module")
	modules.Register("k6/x/xk6_totp", new(TOTP))
}

// TOTP structure
type TOTP struct{}

func (*TOTP) Hello(name string) string {
	return "Hello " + name
}
