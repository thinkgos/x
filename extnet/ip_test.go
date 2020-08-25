package extnet

import (
	"net"
	"reflect"
	"testing"
)

func TestToNumeric(t *testing.T) {
	tests := []struct {
		name    string
		ip      net.IP
		want    uint32
		wantErr bool
	}{
		{"valid IP", net.ParseIP("10.10.0.1"), 0x0a0a0001, false},
		{"invalid IP", net.ParseIP("10.10.x"), 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToNumeric(tt.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNumeric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPv4(t *testing.T) {
	tests := []struct {
		name string
		v    uint32
		want net.IP
	}{
		{"numb", 0x0a0a0001, net.ParseIP("10.10.0.1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPv4(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNumeric(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    uint32
		wantErr bool
	}{
		{"ip string to number", "10.10.0.1", 0x0a0a0001, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNumeric(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNumeric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
