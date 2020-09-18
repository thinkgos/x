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
	"math"
	"net"
)

const (
	IPv4Uint32Cnt = 1
	IPv6Uint32Cnt = 4
)

// IPv4 数值转换为net.ToIP
func ToIP(v uint32) net.IP {
	return net.IPv4(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// Numeric represents an ToIP address using uint32 as internal storage.
// IPv4 uses 1 uint32
// IPv6 uses 4 uint32.
type Numeric []uint32

// IPv4 returns a equivalent Numeric to given uint32 number,
func IPv4(a uint32) Numeric {
	return []uint32{a}
}

// IPv4 returns a equivalent Numeric to given uint32 number,
func IPv6(a, b, c, d uint32) Numeric {
	return []uint32{a, b, c, d}
}

// IP returns a equivalent Numeric to given IP address,
// return nil if ip is neither IPv4 nor IPv6.
func IP(ip net.IP) Numeric {
	if ip == nil {
		return nil
	}
	coercedIP, parts := ip.To4(), 1
	if coercedIP == nil {
		coercedIP, parts = ip.To16(), 4
		if coercedIP == nil {
			return nil
		}
	}

	nn := make(Numeric, parts)
	for i := 0; i < parts; i++ {
		idx := i * net.IPv4len
		nn[i] = binary.BigEndian.Uint32(coercedIP[idx : idx+net.IPv4len])
	}
	return nn
}

// ToIP returns equivalent net.IP.
func (n Numeric) ToIP() net.IP {
	if len(n) != 1 && len(n) != 4 {
		return nil
	}
	ip := make(net.IP, len(n)*net.IPv4len)
	for i := 0; i < len(n); i++ {
		idx := i * net.IPv4len
		binary.BigEndian.PutUint32(ip[idx:idx+net.IPv4len], n[i])
	}
	if len(ip) == net.IPv4len {
		ip = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	}
	return ip
}

// To4 returns ip address if ip is IPv4, returns nil otherwise.
func (n Numeric) To4() Numeric {
	if len(n) != IPv4Uint32Cnt {
		return nil
	}
	return n
}

// To16 returns ip address if ip is IPv6, returns nil otherwise.
func (n Numeric) To16() Numeric {
	if len(n) != IPv6Uint32Cnt {
		return nil
	}
	return n
}

// Equal is the equality test for 2 network numbers.
func (n Numeric) Equal(n1 Numeric) bool {
	if len(n) != len(n1) {
		return false
	}
	if n[0] != n1[0] {
		return false
	}
	if len(n) == IPv6Uint32Cnt {
		return n[1] == n1[1] && n[2] == n1[2] && n[3] == n1[3]
	}
	return true
}

// Previous returns the previous logical network number.
func (n Numeric) Previous() Numeric {
	newIP := make(Numeric, len(n))
	copy(newIP, n)
	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]--
		if newIP[i] < math.MaxUint32 {
			break
		}
	}
	return newIP
}

// Next returns the next logical network number.
func (n Numeric) Next() Numeric {
	newIP := make(Numeric, len(n))
	copy(newIP, n)
	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]++
		if newIP[i] > 0 {
			break
		}
	}
	return newIP
}
