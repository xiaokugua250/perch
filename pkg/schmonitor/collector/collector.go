package collector

import "perch/pkg/schmonitor"

type Collector interface {
	CollectorRegisterWithOpt(server schmonitor.ServerTarget) error
	CollectorAgent() error
	CollectorConnection(server schmonitor.ServerTarget) error
	//CollectorPush() error
}

type LoggerCollector struct {
	LogLocation string
}

type FileCollector struct {
}

type GenCollector struct {
}

type BasicCollector struct {
}

func (log *LoggerCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
	case schmonitor.HttpProtocol:
	case schmonitor.GrpcProtocol:
	case schmonitor.SocketProtocol:
	default:
		return nil

	}
	return err
}

func (log *LoggerCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}

/**
日志采集
*/
func (log *LoggerCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}

//----
func (file *FileCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
	case schmonitor.HttpProtocol:
	case schmonitor.GrpcProtocol:
	case schmonitor.SocketProtocol:
	default:
		return nil

	}
	return err

}

/**
文件采集
*/
func (file *FileCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}
func (file *FileCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}

//--

func (basic *BasicCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
	case schmonitor.HttpProtocol:
	case schmonitor.GrpcProtocol:
	case schmonitor.SocketProtocol:
	default:
		return nil

	}
	return err
}

func (basic *BasicCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}

func (basic *BasicCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}
