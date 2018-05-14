// +build !windows

package globalenv

import (
	"os/exec"
	"strings"
)

func (h Handler) supported() bool {
	_, err := h.execEnv()
	return err == nil
}

func (h Handler) environ() []string {
	out, err := h.execEnv()
	if err != nil {
		return nil
	}

	var env []string
	for _, line := range strings.Split(string(out), "\n") {
		if line == "" {
			continue
		}

		env = append(env, line)
	}

	return env
}

func (h Handler) execEnv() ([]byte, error) {
	cmd := exec.Command(
		"/usr/bin/env", "-", h.FallbackShell,
		"-c", "if test -f /etc/profile ; then . /etc/profile ; fi ; env")
	return cmd.CombinedOutput()
}
