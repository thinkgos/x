package extnet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsDomain(t *testing.T) {
	tests := []struct {
		name string
		host string
		want bool
	}{
		{
			"domain",
			"localhost",
			true,
		},
		{
			"ip",
			"127.0.0.1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDomain(tt.host); got != tt.want {
				t.Errorf("IsDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsIntranet(t *testing.T) {
	tests := []struct {
		name string
		host string
		want bool
	}{
		{
			"ipv4 Loopback 127.0.0.0~127.255.255.255",
			"127.1.1.1",
			true,
		},
		{
			"ipv4 Loopback localhost",
			"localhost",
			true,
		},
		{
			"ipv6 Loopback",
			net.IPv6loopback.String(),
			true,
		},

		{
			"A类10.0.0.0~10.255.255.255",
			"10.1.1.1",
			true,
		},
		{
			"not in A类10.0.0.0~10.255.255.255",
			"11.1.1.1",
			false,
		},
		{
			"b类172.16.0.0~172.31.255.255",
			"172.16.1.1",
			true,
		},
		{
			"1 - not in b类172.16.0.0~172.31.255.255 ",
			"172.15.1.1",
			false,
		},
		{
			"2 - not in b类172.16.0.0~172.31.255.255",
			"172.32.1.1",
			false,
		},
		{
			"c类192.168.0.0~192.168.255.255",
			"192.168.1.1",
			true,
		},
		{
			"not in c类192.168.0.0~192.168.255.255",
			"192.169.1.1",
			false,
		},
		{
			"not intranet",
			"www.baidu.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntranet(tt.host); got != tt.want {
				t.Errorf("IsIntranet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsIntranet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsIntranet("192.168.1.1")
	}
}

func TestSplitHostPort(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		want    string
		want1   uint16
		wantErr bool
	}{
		{
			"localhost:8080",
			"localhost:8080",
			"localhost",
			8080,
			false,
		},
		{
			"127.0.0.1:8080",
			"127.0.0.1:8080",
			"127.0.0.1",
			8080,
			false,
		},

		{
			"[::1]:8080",
			"[::1]:8080",
			"::1",
			8080,
			false,
		},
		{
			"[::1%lo0]:8080",
			"[::1%lo0]:8080",
			"::1%lo0",
			8080,
			false,
		},
		{
			"invalid addr",
			"127.0.0.1",
			"",
			0,
			true,
		},
		{
			"invalid addr port",
			"127.0.0.1:a",
			"",
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := SplitHostPort(tt.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitHostPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SplitHostPort() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitHostPort() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestJoinHostPort(t *testing.T) {
	require.Equal(t, "localhost:8080", JoinHostPort("localhost", 8080))
}
