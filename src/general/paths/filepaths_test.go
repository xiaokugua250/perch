package paths

import (
	"fmt"
	"testing"
)

func TestGetExecableProgramRealPath(t *testing.T) {
	err, path := GetExecableProgramRealPath()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("real path is:", path)
}
