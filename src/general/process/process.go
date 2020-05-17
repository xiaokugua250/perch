package process

import (
	"os"
)

func GetCurrentProcessID() int {

	return os.Getpid()
}

func GetParentProcessID() int {
	return os.Getppid()
}
