package globalenv

import (
	"runtime"
	"strings"
)

var DefaultShell = "/bin/sh"

type Handler struct {
	// FallbackShell is used on UNIX-like systems to source the environment.
	// If it is not available, globalenv will not work on Linux and macOS.
	FallbackShell string
	// SystemWide indices if only system-wide environment variables should be
	// considered. If false, user-specific environment will be retrieved.
	SystemWide bool
}

// DefaultHandler is used for global functions in this package.
var DefaultHandler = Handler{
	FallbackShell: "/bin/sh",
	SystemWide:    true,
}

// Supported calls `Supported` on  `DefaultHandler.`
func Supported() bool {
	return DefaultHandler.Supported()
}

// Environ calls `Environ` on `DefaultHandler`.
func Environ() []string {
	return DefaultHandler.Environ()
}

// Getenv calls `Getenv` on `DefaultHandler`.
func Getenv(key string) string {
	return DefaultHandler.Getenv(key)
}

// Supported returns true if this handler is able to get the environment in the
// current system. If it returns false, all functions will return error or zero
// values.
func (h Handler) Supported() bool {
	return h.supported()
}

// Environ returns all environment variables.
func (h Handler) Environ() []string {
	return h.environ()
}

// Getenv returns the value of the given environment variable.
// It returns an empty string if the variable is not found.
func (h Handler) Getenv(key string) string {
	env := h.Environ()
	for _, line := range env {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}

		if runtime.GOOS == "windows" {
			if strings.ToLower(kv[0]) == strings.ToLower(key) {
				return kv[1]
			}
		} else {
			if kv[0] == key {
				return kv[1]
			}
		}

	}

	return ""
}
