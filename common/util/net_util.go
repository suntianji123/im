package util

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"net"
)

var NetUtil = &netUtil{}

func init() {
	NetUtil.ipv4 = NetUtil.genIpv4()
}

type netUtil struct {
	ipv4 string
}

func (u *netUtil) genIpv4() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		logger.Errorf("get current host ip err: ", err)
		return "127.0.0.1"
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}

func (u *netUtil) Ipv4() string {
	return u.ipv4
}
