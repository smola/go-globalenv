package globalenv_test

import (
	"fmt"

	"github.com/smola/go-globalenv"
)

func Example() {
	path := globalenv.Getenv("PATH")
	if len(path) > 0 {
		fmt.Println("PATH is globally set")
	}
	// Output: PATH is globally set
}
