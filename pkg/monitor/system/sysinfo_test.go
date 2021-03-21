/**
利用gopsutils 包获取服务器信息
ref https://github.com/shirou/gopsutil
*/
package system

import (
	"fmt"
	"github.com/shirou/gopsutil/docker"
	"testing"
)

func TestSysCommonInfo(t *testing.T) {
	meminfo, err := SysMemInfo()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%#v\n", meminfo.SwapMem.Total)
}

func TestSysHostAdvancedInfo(t *testing.T) {
	host, err := SysHostAdvancedInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", host)
}

func TestSysAdvancedDockerInfo(t *testing.T) {
	list, err := docker.GetDockerIDList()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range list {
		fmt.Println(item)
	}
	/*docker,err := SysAdvancedDockerInfo()
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Printf("%#v",docker)*/
}
