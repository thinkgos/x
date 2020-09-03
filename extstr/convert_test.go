package extstr

import (
	"reflect"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"string to bytes no copy", args{"hello world"}, []byte("hello world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Bytes(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2Str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"bytes to string no copy", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2Str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var testBytes = []byte(testString)

func BenchmarkStr2Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes(testString)
	}
}

func BenchmarkStr2BytesCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testString)
	}
}

func BenchmarkBytes2Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes2Str(testBytes)
	}
}

func BenchmarkBytes2StrCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBytes)
	}
}
