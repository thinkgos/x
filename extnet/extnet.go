// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extnet extend network
package extnet

import (
	"net"
	"strconv"
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

// SplitHostPort splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
// host%zone and port.
//
// A literal IPv6 address in hostport must be enclosed in square
// brackets, as in "[::1]:80", "[::1%lo0]:80".
//
// See func Dial for a description of the hostport parameter, and host
// and port results.
func SplitHostPort(addr string) (string, uint16, error) {
	host, p, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	port, err := strconv.ParseUint(p, 10, 16)
	if err != nil {
		return "", 0, err
	}
	return host, uint16(port), nil
}

// JoinHostPort combines host and port into a network address of the
// form "host:port". If host contains a colon, as found in literal
// IPv6 addresses, then JoinHostPort returns "[host]:port".
//
// See func Dial for a description of the host and port parameters.
func JoinHostPort(host string, port uint16) string {
	return net.JoinHostPort(host, strconv.FormatUint(uint64(port), 10))
}
