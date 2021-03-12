/**
利用gopsutils 包获取服务器信息
ref https://github.com/shirou/gopsutil
*/
package system

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"runtime"
	"time"
)

type SysMemInformation struct {
	SwapMem    mem.SwapMemoryStat    `json:"swap_mem"`
	VirtaulMem mem.VirtualMemoryStat `json:"virtaul_mem"`
}

type HostAdvancedInfo struct {
	KernelVersion  string                 `json:"kernel_version"`
	Platform       string                 `json:"platform"`
	Family         string                 `json:"family"`
	Version        string                 `json:"version"`
	Vituralization []string               `json:"vituralization"`
	InfoStat       *host.InfoStat         `json:"info_stat"`
	BootTime       uint64                 `json:"boot_time"`
	UpTime         uint64                 `json:"up_time"`
	Temperature    []host.TemperatureStat `json:"temperature"`
	Users          []host.UserStat        `json:"users"`
}
type CpuAdvancedInfo struct {
	CpuInfoStat []cpu.InfoStat  `json:"cpu_info_stat"`
	CpuCounts   int             `json:"cpu_counts"`
	Percents    []float64       `json:"percents"`
	TImeStat    []cpu.TimesStat `json:"t_ime_stat"`
}
type DiskAdvacedInfo struct {
	SerialNumber string                         `json:"serial_number"`
	Lables       string                         `json:"lables"`
	IOCounters   map[string]disk.IOCountersStat `json:"io_counters"`
	Partitions   []disk.PartitionStat           `json:"partitions"`
	Usage        *disk.UsageStat                `json:"usage"`
}

type DockerAdvancedInfo struct {
	DockerIds   []string                  `json:"docker_ids"`
	DockerStats []docker.CgroupDockerStat `json:"docker_stats"`
}

type LoadAdvancedInfo struct {
	AvgStat  *load.AvgStat `json:"avg_stat"`
	MiscStat *load.MiscStat
}
type NetAdvancedInfo struct {
	Pids            []int32             `json:"pids"`
	ConnectionStats []net.ConntrackStat `json:"connection_stats"`
	InterfacesStat  []net.InterfaceStat
}

type ProcessAdvancedInfo struct {
	Pids      []int32            `json:"pids"`
	Processes []*process.Process `json:"processes"`
}

func init() {
	if runtime.GOOS == "windows" {
		//todo
	}
}

//硬盘
func SysAdvancedDiskInfo(diskSerialName string, diskLableName string, partions bool, path string, iocounters ...string) (DiskAdvacedInfo, error) {
	var (
		diskAdvancedInfo DiskAdvacedInfo
		err              error
	)
	if runtime.GOOS == "windows" {
		return diskAdvancedInfo, errors.New("current system is windows!,only linux support!")
	}
	if diskSerialName != "" {
		//	diskAdvancedInfo.SerialNumber= disk.GetDiskSerialNumber(diskSerialName)
	}

	if diskLableName != "" {
		//	diskAdvancedInfo.Lables=disk.GetLabel(diskLableName)
	}
	if len(iocounters) >= 1 {
		diskAdvancedInfo.IOCounters, err = disk.IOCounters(iocounters...)
		if err != nil {
			return DiskAdvacedInfo{}, err
		}
	}

	diskAdvancedInfo.Partitions, err = disk.Partitions(partions)
	if err != nil {
		return DiskAdvacedInfo{}, err
	}
	if path != "" {
		diskAdvancedInfo.Usage, err = disk.Usage(path)
		if err != nil {
			return DiskAdvacedInfo{}, err
		}
	}

	return diskAdvancedInfo, err

}
func SysAdvancedCpuInfo(logical, percpu bool, interval time.Duration) (CpuAdvancedInfo, error) {
	var (
		cpuAdvancedInfo CpuAdvancedInfo
		err             error
	)
	cpuAdvancedInfo.CpuInfoStat, err = cpu.Info()
	if err != nil {
		return CpuAdvancedInfo{}, err
	}
	cpuAdvancedInfo.TImeStat, err = cpu.Times(percpu)
	if err != nil {
		return CpuAdvancedInfo{}, err
	}
	cpuAdvancedInfo.CpuCounts, err = cpu.Counts(logical)
	if err != nil {
		return CpuAdvancedInfo{}, err
	}
	cpuAdvancedInfo.Percents, err = cpu.Percent(interval, percpu)
	if err != nil {
		return CpuAdvancedInfo{}, err
	}

	return cpuAdvancedInfo, err

}

//系统负载
func SysAdvancedLoadInfo() (LoadAdvancedInfo, error) {
	var (
		loadInfo LoadAdvancedInfo
		err      error
	)
	loadInfo.AvgStat, err = load.Avg()
	if err != nil {
		return LoadAdvancedInfo{}, err
	}
	loadInfo.MiscStat, err = load.Misc()
	if err != nil {
		return LoadAdvancedInfo{}, err
	}
	return loadInfo, err

}

func SysAdvancedDockerInfo() (DockerAdvancedInfo, error) {
	var (
		dockerinfo DockerAdvancedInfo
		err        error
	)
	dockerinfo.DockerIds, err = docker.GetDockerIDList()
	if err != nil {
		return DockerAdvancedInfo{}, err
	}
	dockerinfo.DockerStats, err = docker.GetDockerStat()
	if err != nil {
		return DockerAdvancedInfo{}, err
	}
	return dockerinfo, err
}
func SysAdvancedNetInfo(percpu bool) (NetAdvancedInfo, error) {
	var (
		netinfo NetAdvancedInfo
		err     error
	)
	if runtime.GOOS == "windows" {
		return netinfo, errors.New("current system is windows!,only linux support!")
	}
	//netinfo.Pids,err = net.Pids()
	if err != nil {
		return NetAdvancedInfo{}, err
	}
	netinfo.ConnectionStats, err = net.ConntrackStats(percpu)
	if err != nil {
		return NetAdvancedInfo{}, err
	}
	netinfo.InterfacesStat, err = net.Interfaces()
	if err != nil {
		return NetAdvancedInfo{}, err
	}

	return netinfo, err
}
func SysAdvancedProcessInfo() (ProcessAdvancedInfo, error) {
	var (
		processinfo ProcessAdvancedInfo
		err         error
	)
	processinfo.Pids, err = process.Pids()
	if err != nil {
		return ProcessAdvancedInfo{}, err
	}
	processinfo.Processes, err = process.Processes()
	if err != nil {
		return ProcessAdvancedInfo{}, err
	}
	return processinfo, err
}

func SysCpuInfo() ([]cpu.InfoStat, error) {

	return cpu.Info()
}

func SysCpuTimeStateInfo(percpu bool) ([]cpu.TimesStat, error) {
	return cpu.Times(percpu)

}
func SysCpuCount(logical bool) (int, error) {
	return cpu.Counts(logical)
}

func SysCpuPercent(interval time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(interval, percpu)
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

func SysWinServiecInfo() {

}

func SysHostAdvancedInfo() (HostAdvancedInfo, error) {
	var (
		advancedInfo HostAdvancedInfo
		err          error
	)
	advancedInfo.KernelVersion, err = host.KernelVersion()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.BootTime, err = host.BootTime()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.UpTime, err = host.Uptime()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.Platform, advancedInfo.Family, advancedInfo.Version, err = host.PlatformInformation()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.InfoStat, err = host.Info()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.Users, err = host.Users()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	advancedInfo.Temperature, err = host.SensorsTemperatures()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	var virtualizationArry []string
	a, b, err := host.Virtualization()
	if err != nil {
		return HostAdvancedInfo{}, err
	}
	virtualizationArry = append(virtualizationArry, a, b)
	advancedInfo.Vituralization = virtualizationArry

	return advancedInfo, err
}

func SysHostPlatformInfo() (platform string, family string, version string, err error) {
	return host.PlatformInformation()
}

func SysHostVirtualizationInfo() (string, string, error) {
	return host.Virtualization()
}

func SysHostStatInfo() (host.InfoStat, error) {
	info, err := host.Info()
	return *info, err

}

func SysHostTemperatureStatInfo() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}

func SysHostUserInfo() ([]host.UserStat, error) {
	return host.Users()
}
func SysHostKernelVersionInfo() (version string, err error) {
	return host.KernelVersion()
}
