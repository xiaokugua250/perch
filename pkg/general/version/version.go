package version

import "runtime"

func GetGoLangRuntimeVersion() (version string) {

	return runtime.Version()
}
