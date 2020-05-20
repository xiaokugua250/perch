package version

import (
	"fmt"
	"testing"
)

func TestGetGoLangRuntimeVersion(t *testing.T) {
	version := GetGoLangRuntimeVersion()
	fmt.Println(version)

}
