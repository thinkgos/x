package extnet

import (
	"net"
)

// IsDomain 是否是域名,只检查host或ip,不可带port,否则会误判
func IsDomain(host string) bool {
	return net.ParseIP(host) == nil
}

// IsIntranet is intranet network,if host is domain,it will looks up host using the local resolver.
// net.LookupIP may cause deadlock in windows
// see https://github.com/golang/go/issues/24178
// 局域网IP段:
// 		A类: 10.0.0.0~10.255.255.255
// 		B类: 172.16.0.0~172.31.255.255
// 		C类: 192.168.0.0~192.168.255.255
func IsIntranet(host string) bool {
	var ips []net.IP
	var err error

	if _ip := net.ParseIP(host); _ip != nil { // is ip
		ips = []net.IP{_ip}
	} else if ips, err = net.LookupIP(host); err != nil { // is domain
		return false
	}

	for _, ip := range ips {
		if ip4 := ip.To4(); ip4 != nil &&
			(ip4[0] == 127 || // ipv4 loopback
				ip4[0] == 10 || // A类
				(ip4[0] == 172 && (ip4[1] >= 16) && (ip4[1] <= 31)) || // B类
				(ip4[0] == 192 && ip4[1] == 168)) || // C类
			ip.Equal(net.IPv6loopback) { // ipv6 loopback
			return true
		}
	}
	return false
}
