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

package extnet

import (
	"encoding/binary"
	"net"
)

// IPv4 数值转换为net.IP
func IPv4(v uint32) net.IP {
	return net.IPv4(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// ToNumeric net.IP转换为数值
func ToNumeric(p net.IP) (uint32, error) {
	if ip := p.To4(); ip != nil {
		return binary.BigEndian.Uint32(ip), nil
	}
	return 0, &net.ParseError{
		Type: "IP address",
		Text: "not ipv4",
	}
}

// ParseNumeric 点分十进制字符串转换数值
func ParseNumeric(s string) (uint32, error) {
	return ToNumeric(net.ParseIP(s))
}
