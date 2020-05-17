package paths

import (
	"os"
	"path/filepath"
)

func GetExecableProgramRealPath() (error, string) {

	execuable, err := os.Executable()
	if err != nil {
		return err, ""
	}
	execuablePath := filepath.Dir(execuable)

	realPath, err := filepath.EvalSymlinks(execuablePath)
	return err, realPath
}
