package network

import "net"

/**
获取设备IP
*/
func GetDeviceIPAddress() (map[string][]string, error) {

	var DeviceIps = make(map[string][]string) //一块网卡，多个IP
	netinterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, interf := range netinterfaces {
		var ips []string
		addrs, err := interf.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {

			if ip, ok := addr.(*net.IPNet); ok {
				ips = append(ips, ip.String())
			}

		}
		DeviceIps[interf.Name] = ips
	}
	return DeviceIps, err
}

/**
创建TCP SERVER
*/
func CreateTCPServer(serverAddress string) (net.Listener, error) {

	server, err := net.Listen("tcp", serverAddress)
	return server, err
}

/**
创建UDP SERVER
*/
func CreateUDPServer(serverAddress string) (net.PacketConn, error) {

	server, err := net.ListenPacket("udp", serverAddress)
	return server, err
}
