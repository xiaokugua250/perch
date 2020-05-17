package env

import "os"

/*
GET OS ENV
*/
func GetOSEnv(envName string, defaultValue string) string {
	value, ok := os.LookupEnv(envName)
	if ok {
		return value
	}
	return defaultValue

}

/**
set env Value
*/
func SetOSenv(envName string, envValue string) error {
	return os.Setenv(envName, envValue)
}
