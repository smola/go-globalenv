// +build windows

package globalenv

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func supported() bool {
	return true
}

func environ() []string {
	envMap, _ := getEnv(registry.LOCAL_MACHINE, `System\CurrentControlSet\Control\Session Manager\Environment`)

	cuEnvMap, _ := getEnv(registry.CURRENT_USER, `Environment`)

	for k, v := range cuEnvMap {
		envMap[k] = v
	}

	var env []string
	for k, v := range envMap {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env

}

func getEnv(key registry.Key, path string) (map[string]string, error) {
	k, err := registry.OpenKey(key, path, registry.READ)
	if err != nil {
		return nil, err
	}

	defer k.Close()
	valueNames, err := k.ReadValueNames(-1)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string, len(valueNames))
	for _, valueName := range valueNames {
		value, _, err := k.GetStringValue(valueName)
		if err != nil {
			continue
		}

		env[valueName] = value
	}

	return env, nil
}
