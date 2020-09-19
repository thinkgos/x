package extnet

import (
	"math"
	"net"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToIP(t *testing.T) {
	tests := []struct {
		name string
		v    uint32
		want net.IP
	}{
		{"numb", 0x0a0a0001, net.ParseIP("10.10.0.1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToIP(tt.v))
		})
	}
}

func TestNextIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		next string
	}{
		{"IPv4 basic", "0.0.0.0", "0.0.0.1"},
		{"IPv4 rollover", "0.0.0.255", "0.0.1.0"},
		{"IPv4 consecutive rollover", "0.255.255.255", "1.0.0.0"},
		{"IPv6 basic", "8000::0", "8000::1"},
		{"IPv6 rollover", "0::ffff", "0::1:0"},
		{"IPv6 consecutive rollover", "0:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "1::"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, net.ParseIP(tt.next), NextIP(net.ParseIP(tt.ip)))
		})
	}
}

func TestPreviousIP(t *testing.T) {
	tests := []struct {
		ip   string
		next string
		name string
	}{
		{"0.0.0.1", "0.0.0.0", "IPv4 basic"},
		{"0.0.1.0", "0.0.0.255", "IPv4 rollover"},
		{"1.0.0.0", "0.255.255.255", "IPv4 consecutive rollover"},
		{"8000::1", "8000::0", "IPv6 basic"},
		{"0::1:0", "0::ffff", "IPv6 rollover"},
		{"1::0", "0:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "IPv6 consecutive rollover"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, net.ParseIP(tt.next), PreviousIP(net.ParseIP(tt.ip)))
		})
	}
}

func TestIPv4(t *testing.T) {
	tests := []struct {
		name string
		a    uint32
		want Numeric
	}{
		{"IPv4", 2130706433, IP(net.ParseIP("127.0.0.1"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IPv4(tt.a))
		})
	}
}

func TestIPv6(t *testing.T) {
	type args struct {
		a uint32
		b uint32
		c uint32
		d uint32
	}
	tests := []struct {
		name string
		args args
		want Numeric
	}{
		{
			"IPv6",
			args{
				536939960, 0, 65280, 4358953,
			},
			IP(net.ParseIP("2001:0db8::ff00:0042:8329")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPv6(tt.args.a, tt.args.b, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP(t *testing.T) {
	tests := []struct {
		name string
		ip   net.IP
		want Numeric
	}{
		{"bad", net.IP([]byte{1, 1, 1, 1, 1}), nil},
		{"nil", nil, nil},
		{"IPv4", net.ParseIP("127.0.0.1"), Numeric([]uint32{2130706433})},
		{
			"IPv6",
			net.ParseIP("2001:0db8::ff00:0042:8329"),
			Numeric([]uint32{536939960, 0, 65280, 4358953}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IP(tt.ip))
		})
	}
}

func TestNumeric_String(t *testing.T) {
	n := Numeric([]uint32{2130706433})
	require.Equal(t, "127.0.0.1", n.String())
}

func TestNumeric_To4_To16(t *testing.T) {
	tests := []struct {
		name string
		ip   Numeric
		to4  Numeric
		to6  Numeric
		toip net.IP
	}{
		{"IPv4", Numeric([]uint32{1}), Numeric([]uint32{1}), nil, net.IPv4(0, 0, 0, 1)},
		{"IPv6", Numeric([]uint32{1, 1, 1, 1}), nil, Numeric([]uint32{1, 1, 1, 1}), net.IP{0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1}},
		{"invalid", Numeric([]uint32{1, 1}), nil, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.to4, tt.ip.To4())
			assert.Equal(t, tt.to6, tt.ip.To16())
			assert.Equal(t, tt.toip, tt.ip.ToIP())
		})
	}
}

func TestNumeric_Equal(t *testing.T) {
	tests := []struct {
		name string
		n    Numeric
		n1   Numeric
		want bool
	}{
		{"IPv4 equals", Numeric{math.MaxUint32}, Numeric{math.MaxUint32}, true},
		{"IPv4 does not equal", Numeric{math.MaxUint32}, Numeric{math.MaxUint32 - 1}, false},
		{"IPv6 equals", Numeric{1, 1, 1, 1}, Numeric{1, 1, 1, 1}, true},
		{"IPv6 does not equal", Numeric{1, 1, 1, 1}, Numeric{1, 1, 1, 2}, false},
		{"Version mismatch", Numeric{1}, Numeric{1, 2, 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.n.Equal(tt.n1))
		})
	}
}

func TestNumeric_Previous(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		previous string
	}{
		{"IPv4 basic", "0.0.0.1", "0.0.0.0"},
		{"IPv4 rollover", "0.0.1.0", "0.0.0.255"},
		{"IPv4 consecutive rollover", "1.0.0.0", "0.255.255.255"},
		{"IPv6 basic", "8000::1", "8000::0"},
		{"IPv6 rollover", "0::1:0", "0::ffff"},
		{"IPv6 consecutive rollover", "1::0", "0:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ip := IP(net.ParseIP(tc.ip))
			expected := IP(net.ParseIP(tc.previous))
			assert.Equal(t, expected, ip.Previous())
		})
	}
}

func TestNumeric_Next(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		next string
	}{
		{"IPv4 basic", "0.0.0.0", "0.0.0.1"},
		{"IPv4 rollover", "0.0.0.255", "0.0.1.0"},
		{"IPv4 consecutive rollover", "0.255.255.255", "1.0.0.0"},
		{"IPv6 basic", "8000::0", "8000::1"},
		{"IPv6 rollover", "0::ffff", "0::1:0"},
		{"IPv6 consecutive rollover", "0:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "1::"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ip := IP(net.ParseIP(tc.ip))
			expected := IP(net.ParseIP(tc.next))
			assert.Equal(t, expected, ip.Next())
		})
	}
}

func TestNumeric_Mask(t *testing.T) {
	tests := []struct {
		name string
		mask NumericMask
		ip   Numeric
		want Numeric
	}{
		{"nop IPv4 mask", NumericMask{math.MaxUint32}, Numeric{math.MaxUint32}, Numeric{math.MaxUint32}},
		{"nop IPv4 mask", NumericMask{math.MaxUint32 - math.MaxUint16}, Numeric{math.MaxUint16 + 1}, Numeric{math.MaxUint16 + 1}},
		{"IPv4 masked", NumericMask{math.MaxUint32 - math.MaxUint16}, Numeric{math.MaxUint32}, Numeric{math.MaxUint32 - math.MaxUint16}},
		{"nop IPv6 mask", NumericMask{math.MaxUint32, 0, 0, 0}, Numeric{math.MaxUint32, 0, 0, 0}, Numeric{math.MaxUint32, 0, 0, 0}},
		{"nop IPv6 masked", NumericMask{math.MaxUint32 - math.MaxUint16, 0, 0, 0}, Numeric{math.MaxUint16 + 1, 0, 0, 0}, Numeric{math.MaxUint16 + 1, 0, 0, 0}},
		{"IPv6 masked", NumericMask{math.MaxUint32 - math.MaxUint16, 0, 0, 0}, Numeric{math.MaxUint32, 0, 0, 0}, Numeric{math.MaxUint32 - math.MaxUint16, 0, 0, 0}},
		{"Version mismatch", NumericMask{math.MaxUint32}, Numeric{math.MaxUint32, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ip.Mask(tt.mask))
		})
	}
}

func TestNewIPNet(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("192.128.0.0/24")
	n := NewIPNet(ipNet)

	assert.Equal(t, ipNet, n.IPNet)
	assert.Equal(t, Numeric{3229614080}, n.Number)
	assert.Equal(t, NumericMask{math.MaxUint32 - uint32(math.MaxUint8)}, n.Mask)
}

func TestIPNet_Masked(t *testing.T) {
	tests := []struct {
		IPNet     string
		mask      int
		wantIPNet string
	}{
		{"192.168.0.0/16", 16, "192.168.0.0/16"},
		{"192.168.0.0/16", 14, "192.168.0.0/14"},
		{"192.168.0.0/16", 18, "192.168.0.0/18"},
		{"192.168.0.0/16", 8, "192.0.0.0/8"},
		{"8000::/128", 96, "8000::/96"},
		{"8000::/128", 128, "8000::/128"},
		{"8000::/96", 112, "8000::/112"},
		{"8000:ffff::/96", 16, "8000::/16"},
	}
	for _, tt := range tests {
		_, network, _ := net.ParseCIDR(tt.IPNet)
		got := NewIPNet(network)
		_, expected, _ := net.ParseCIDR(tt.wantIPNet)
		want := NewIPNet(expected)
		assert.Equal(t, want.String(), got.Masked(tt.mask).String())
	}
}

func TestIPNet_ContainsNumeric(t *testing.T) {
	tests := []struct {
		name    string
		ipNet   string
		firstIP string
		lastIP  string
	}{
		{"192.168.0.0/24 contains", "192.168.0.0/24", "192.168.0.0", "192.168.0.255"},
		{"8000::0/120 contains", "8000::0/120", "8000::0", "8000::ff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, net1, _ := net.ParseCIDR(tt.ipNet)
			ipNet := NewIPNet(net1)
			ip := IP(net.ParseIP(tt.firstIP))
			lastIP := IP(net.ParseIP(tt.lastIP))
			assert.False(t, ipNet.ContainsNumeric(ip.Previous()))
			assert.False(t, ipNet.ContainsNumeric(lastIP.Next()))
			for ; !ip.Equal(lastIP.Next()); ip = ip.Next() {
				assert.True(t, ipNet.ContainsNumeric(ip))
			}
		})
	}
}

func TestIPNet_ContainsIPNet(t *testing.T) {
	tests := []struct {
		name   string
		ipNet  string
		covers string
		want   bool
	}{
		{"contains", "10.0.0.0/24", "10.0.0.1/25", true},
		{"not contains", "10.0.0.0/24", "11.0.0.1/25", false},
		{"prefix false", "10.0.0.0/16", "10.0.0.0/15", false},
		{"prefix true", "10.0.0.0/15", "10.0.0.0/16", true},
		{"same", "10.0.0.0/15", "10.0.0.0/15", true},
		{"ip version mismatch", "10::0/15", "10.0.0.0/15", false},
		{"ipv6", "10::0/15", "10::0/16", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, n, _ := net.ParseCIDR(tt.ipNet)
			network := NewIPNet(n)
			_, n, _ = net.ParseCIDR(tt.covers)
			covers := NewIPNet(n)
			assert.Equal(t, tt.want, network.ContainsIPNet(covers))
		})
	}
}

func TestIPNet_Equal(t *testing.T) {
	tests := []struct {
		name string
		n1   string
		n2   string
		want bool
	}{
		{"IPv4 equals", "192.128.0.0/24", "192.128.0.0/24", true},
		{"IPv4 not equals", "192.128.0.0/24", "192.128.0.0/23", false},
		{"IPv6 equals", "8000::/24", "8000::/24", true},
		{"IPv6 not equals", "8000::/24", "8000::/23", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, ipNet1, _ := net.ParseCIDR(tc.n1)
			_, ipNet2, _ := net.ParseCIDR(tc.n2)
			assert.Equal(t, tc.want, NewIPNet(ipNet1).Equal(NewIPNet(ipNet2)))
		})
	}
}
