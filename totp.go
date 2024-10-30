package totp

import (
	"fmt"

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
