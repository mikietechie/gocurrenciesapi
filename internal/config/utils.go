package config

import "os"

func GetEnvOrDef(key, def string) string {
	env := os.Getenv(key)
	if len(env) != 0 {
		return env
	}
	return def
}
