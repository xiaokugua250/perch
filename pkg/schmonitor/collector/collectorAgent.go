package collector

import (
	"github.com/shirou/gopsutil/host"
	"github.com/sony/sonyflake"
	"strconv"
)

type AgentBasicInfo struct {
	ID         int
	UUID       string
	Name       string
	IP         string
	Location   string
	SystemInfo interface{}
}

func AgentExector() error {
	var (
		err error
	)
	return err
}

/**
agent 基本信息采集器
*/
func AgentBasicCollector() (AgentBasicInfo, error) {
	var (
		err       error
		basicInfo AgentBasicInfo
	)
	flake := sonyflake.NewSonyflake(sonyflake.Settings{}) //采用分布式UUID生成算法，避免可能出现重复的问题
	id, err := flake.NextID()
	if err != nil {
		return basicInfo, err
	}
	basicInfo.UUID = strconv.FormatUint(id, 10)
	// Note: this is base16, could shorten by encoding as base62 string
	basicInfo.SystemInfo, err = host.Info()
	//fmt.Printf("++%+v\n",basicInfo)

	return basicInfo, err

}
