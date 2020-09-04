/**
利用gopsutils 包获取服务器信息
ref https://github.com/shirou/gopsutil
*/
package sysinfo

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
	"runtime"
	"time"
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
func SysCpuInfo() ([]cpu.InfoStat,error){

	return cpu.Info()
}

func SysCpuTimeStateInfo(percpu bool)([]cpu.TimesStat,error)  {
	return cpu.Times(percpu)

}
func SysCpuCount(logical bool)(int ,error){
	return cpu.Counts(logical)
}

func SysCpuPercent(interval time.Duration,percpu bool)([]float64,error){
	return cpu.Percent(interval,percpu)
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

type HostTimeInfo struct {
	BootTime uint64 `json:"boot_time"`
	UpTime uint64 `json:"up_time"`
}
func SysHostBootTimeInfo()(HostTimeInfo, error)  {
	var (
		hosttime HostTimeInfo
		err error
	)
	hosttime.BootTime,err= host.BootTime()
	if err!= nil{
		return hosttime,err
	}
	hosttime.UpTime ,err= host.Uptime()
	if err!= nil{
		return hosttime,err
	}
	return hosttime,nil

}

func SysHostKernelVersionInfo() (version string, err error){
	return host.KernelVersion()
}

func SysHostPlatformInfo()(platform string, family string, version string, err error){
	return host.PlatformInformation()
}

func SysHostVirtualizationInfo()(string, string, error){
	return host.Virtualization()
}

func SysHostStatInfo()  (host.InfoStat, error) {
	info,err := host.Info()
	return *info,err

}

func SysHostTemperatureStatInfo() ([]host.TemperatureStat, error){
	return host.SensorsTemperatures()
}

func SysHostUserInfo()([]host.UserStat, error){
	return host.Users()
}
