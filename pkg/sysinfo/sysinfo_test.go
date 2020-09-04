/**
利用gopsutils 包获取服务器信息
ref https://github.com/shirou/gopsutil
*/
package sysinfo

import (
	"fmt"
	"testing"
)

func TestSysCommonInfo(t *testing.T) {
	meminfo, err := SysMemInfo()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%#v\n", meminfo.SwapMem.Total)
}
