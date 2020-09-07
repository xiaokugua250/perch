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



type HostAdvancedInfo struct {

	KernelVersion string `json:"kernel_version"`
	Platform string `json:"platform"`
	Family string `json:"family"`
	Version string `json:"version"`
	Vituralization []string `json:"vituralization"`
	InfoStat *host.InfoStat `json:"info_stat"`
	BootTime uint64 `json:"boot_time"`
	UpTime uint64 `json:"up_time"`
	Temperature []host.TemperatureStat `json:"temperature"`
	Users []host.UserStat `json:"users"`
}
type CpuAdvancedInfo struct {

}
type DiskAdvacedInfo struct {

}

type DockerAdvancedInfo struct {

}

type LoadAdvancedInfo struct {

}
type NetAdvancedInfo struct {

}

type ProcessAdvancedInfo struct {

}
func init() {
	if runtime.GOOS == "windows" {
		//todo
	}
}

func SysAdvancedCpuInfo(logical,percpu bool){

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

func SysHostAdvancedInfo() (HostAdvancedInfo,error) {
	var (
		advancedInfo HostAdvancedInfo
		err error
	)
	advancedInfo.KernelVersion,err = host.KernelVersion()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.BootTime,err= host.BootTime()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.UpTime ,err= host.Uptime()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.Platform,advancedInfo.Family,advancedInfo.Version,err = host.PlatformInformation()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.InfoStat,err= host.Info()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.Users,err= host.Users()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	advancedInfo.Temperature,err= host.SensorsTemperatures()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	var virtualizationArry []string
	a,b,err:= host.Virtualization()
	if err!= nil{
		return HostAdvancedInfo{},err
	}
	virtualizationArry= append(virtualizationArry,a,b)
	advancedInfo.Vituralization= virtualizationArry

	return advancedInfo,err
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
func SysHostKernelVersionInfo() (version string, err error){
	return host.KernelVersion()
}
