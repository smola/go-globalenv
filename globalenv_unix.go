// +build !windows

package globalenv

import (
	"os/exec"
	"strings"
)

func supported() bool {
	_, err := execEnv()
	return err == nil
}

func environ() []string {
	out, err := execEnv()
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

func execEnv() ([]byte, error) {
	cmd := exec.Command("/usr/bin/env", "-", "/bin/sh", "-c",
		". /etc/profile && env")
	return cmd.CombinedOutput()
}
