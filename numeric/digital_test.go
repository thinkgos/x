// Copyright [2020] [thinkgos]
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

package numeric

import (
	"reflect"
	"testing"
)

func TestBreakUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name       string
		args       args
		wantLoByte byte
		wantHiByte byte
	}{
		{"uint16 break into byte", args{0x1234}, 0x34, 0x12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLoByte, gotHiByte := BreakUint16(tt.args.v)
			if gotLoByte != tt.wantLoByte {
				t.Errorf("BreakUint16() gotLoByte = %v, want %v", gotLoByte, tt.wantLoByte)
			}
			if gotHiByte != tt.wantHiByte {
				t.Errorf("BreakUint16() gotHiByte = %v, want %v", gotHiByte, tt.wantHiByte)
			}
		})
	}
}

func TestBreakUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name      string
		args      args
		wantByte0 byte
		wantByte1 byte
		wantByte2 byte
		wantByte3 byte
	}{
		{"uint32 break into byte", args{0x12345678}, 0x78, 0x56, 0x34, 0x12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotByte0, gotByte1, gotByte2, gotByte3 := BreakUint32(tt.args.v)
			if gotByte0 != tt.wantByte0 {
				t.Errorf("BreakUint32() gotByte0 = %v, want %v", gotByte0, tt.wantByte0)
			}
			if gotByte1 != tt.wantByte1 {
				t.Errorf("BreakUint32() gotByte1 = %v, want %v", gotByte1, tt.wantByte1)
			}
			if gotByte2 != tt.wantByte2 {
				t.Errorf("BreakUint32() gotByte2 = %v, want %v", gotByte2, tt.wantByte2)
			}
			if gotByte3 != tt.wantByte3 {
				t.Errorf("BreakUint32() gotByte3 = %v, want %v", gotByte3, tt.wantByte3)
			}
		})
	}
}

func TestBreakUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name     string
		args     args
		wantLo32 uint32
		wantHi32 uint32
	}{
		{"uint64 break into uint32", args{0x112347890}, 0x12347890, 0x1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLo32, gotHi32 := BreakUint64(tt.args.v)
			if gotLo32 != tt.wantLo32 {
				t.Errorf("BreakUint64() gotLo32 = %v, want %v", gotLo32, tt.wantLo32)
			}
			if gotHi32 != tt.wantHi32 {
				t.Errorf("BreakUint64() gotHi32 = %v, want %v", gotHi32, tt.wantHi32)
			}
		})
	}
}

func TestBuildUint16(t *testing.T) {
	type args struct {
		loByte byte
		Hibyte byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"build byte into uint16", args{0x34, 0x12}, 0x1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUint16(tt.args.loByte, tt.args.Hibyte); got != tt.want {
				t.Errorf("BuildUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUint32(t *testing.T) {
	type args struct {
		Byte0 byte
		Byte1 byte
		Byte2 byte
		Byte3 byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"build byte into uint32", args{0x78, 0x56, 0x34, 0x12}, 0x12345678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUint32(tt.args.Byte0, tt.args.Byte1, tt.args.Byte2, tt.args.Byte3); got != tt.want {
				t.Errorf("BuildUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUint64(t *testing.T) {
	type args struct {
		Lo32 uint32
		Hi32 uint32
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"build uint32 into uint64", args{0x12345678, 0x1}, 0x112345678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUint64(tt.args.Lo32, tt.args.Hi32); got != tt.want {
				t.Errorf("BuildUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBytes(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Reverse bytes", args{[]byte{1, 2, 3, 4, 5}}, []byte{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBytes(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
