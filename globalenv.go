package globalenv

import (
	"fmt"
	"strings"
)

// DefaultShell is used on UNIX-like systems to source the environment.
// If it is not available, globalenv will not work.
var DefaultShell = "/bin/sh"

// Supported returns true if `globalenv` is able to get the environment in the
// current system. If it returns false, all functions will return error or zero
// values.
func Supported() bool {
	return supported()
}

func Environ() []string {
	return environ()
}

func Getenv(key string) string {
	env := Environ()
	for _, line := range env {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}

		if kv[0] == key {
			return kv[1]
		}
	}

	return ""
}

func Setenv(key, value string) error {
	return fmt.Errorf("Setenv not supported")
}
