/**
利用gopsutils 包获取服务器信息
ref https://github.com/shirou/gopsutil
*/
package sysinfo

import (
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

type SysMemInformation struct {
	SwapMem    mem.SwapMemoryStat    `json:"swap_mem"`
	VirtaulMem mem.VirtualMemoryStat `json:"virtaul_mem"`
}

func init() {
	if runtime.GOOS == "windows" {
		//todo
	}
}
func SysCpuInfo() {

}

//硬盘
func SysDiskInfo() {

}

func SysDockerInfo() {

}

func SysCommonInfo() {

}

//系统负载
func SysLoadInfo() {

}

func SysMemInfo() (SysMemInformation, error) {
	var (
		MemInformation SysMemInformation
		virtualMem     *mem.VirtualMemoryStat
		swapMem        *mem.SwapMemoryStat
		err            error
	)

	virtualMem, err = mem.VirtualMemory()
	if err != nil {
		return MemInformation, err
	}
	swapMem, err = mem.SwapMemory()
	if err != nil {
		return MemInformation, err

	}

	MemInformation.SwapMem = *swapMem
	MemInformation.VirtaulMem = *virtualMem

	return MemInformation, nil

}

func SysNetInfo() {

}

func SysProcessInfo() {

}

func SysWinServiecInfo() {

}
