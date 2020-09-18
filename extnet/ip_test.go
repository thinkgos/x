package extnet

import (
	"math"
	"net"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestNumeric_To4_To16(t *testing.T) {
	cases := []struct {
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
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.to4, tc.ip.To4())
			assert.Equal(t, tc.to6, tc.ip.To16())
			assert.Equal(t, tc.toip, tc.ip.ToIP())
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
