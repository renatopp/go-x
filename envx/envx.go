package envx

import (
	"os"
	"strings"
)

// Has checks if the environment variable with the given key is set.
func Has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// Get returns the value of the environment variable with the given key.
// If the variable is not set, it returns an empty string.
func Get(key string) string {
	return os.Getenv(key)
}

// GetOr returns the value of the environment variable with the given key
// if it is set; otherwise, it returns the provided default value.
func GetOr(key, value string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return value
}

// GetOk returns the value of the environment variable with the given key
// and a boolean indicating whether the variable is set.
func GetOk(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Set sets the value of the environment variable with the given key.
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// Unset unsets the environment variable with the given key.
func Unset(key string) error {
	return os.Unsetenv(key)
}

// Clean clears all environment variables, effectively resetting the
// environment to an empty state.
func Clean() {
	os.Clearenv()
}

// Expand expands environment variables in the given string using the current
// environment. It replaces ${var} or $var in the string according to the values
// of the current environment variables.
func Expand(value string) string {
	return os.ExpandEnv(value)
}

// List returns a slice of all environment variables in the form "key=value".
func List() []string {
	return os.Environ()
}

// Keys returns a slice of all environment variable keys.
func Keys() []string {
	envs := os.Environ()
	keys := make([]string, len(envs))
	for i, env := range envs {
		keys[i] = env[:strings.Index(env, "=")]
	}
	return keys
}

// Map returns a map of all environment variables with their keys and values.
func Map() map[string]string {
	envs := os.Environ()
	m := make(map[string]string, len(envs))
	for _, env := range envs {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			m[parts[0]] = parts[1]
		}
	}
	return m
}
