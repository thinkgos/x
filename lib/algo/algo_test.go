package algo

import (
	"testing"
)

func TestMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{s: "thinkgos"},
			want: "f2cd401856d28a5e3b6bf963416229a1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sha1",
			args: args{s: "thinkgos"},
			want: "a36018b908e4add1e7e993599cc0cf9f26025c1f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		method Method
		s      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid method",
			args: args{method: "", s: "thinkgos"},
			want: "thinkgos",
		},
		{
			name: "md5",
			args: args{method: MethodMD5, s: "thinkgos"},
			want: "f2cd401856d28a5e3b6bf963416229a1",
		},
		{
			name: "sha1",
			args: args{method: MethodSha1, s: "thinkgos"},
			want: "a36018b908e4add1e7e993599cc0cf9f26025c1f",
		},
		{
			name: "sha224",
			args: args{method: MethodSha224, s: "thinkgos"},
			want: "1d110ce70effcee182104b635b0edddfd541c62f0cbcb8c2405e3fb5",
		},
		{
			name: "sha256",
			args: args{method: MethodSha256, s: "thinkgos"},
			want: "75b95632b997d60ab858d58196d3a07196d280e7ed10c8eaeebebc8ef5be2ec4",
		},
		{
			name: "sha384",
			args: args{method: MethodSha384, s: "thinkgos"},
			want: "0c9766ebf8d19a48584566c0df0bd714bc8319c45786c5934984dd4cdabe3000c4eef1a6ac74fabcc6e36229db351c8b",
		},
		{
			name: "sha512",
			args: args{method: MethodSha512, s: "thinkgos"},
			want: "11fae529a07be46ed230cb0bac0e66c9b1a31b5a1a745c7048c31579cd55c5e757d9d90967117c08ed989baae99851b1dea54c338a7d9ed4b40dccb0a5aa1980",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.method, tt.args.s); got != tt.want {
				t.Errorf("Hash(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestHmac(t *testing.T) {
	type args struct {
		method Method
		s      string
		key    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid method",
			args: args{method: "", s: "thinkgos", key: "thinkgos"},
			want: "thinkgos",
		},
		{
			name: "md5",
			args: args{method: MethodMD5, s: "thinkgos", key: "thinkgos"},
			want: "ed9dc7baf84a9740a3ceb6f1f26bfb4f",
		},
		{
			name: "sha1",
			args: args{method: MethodSha1, s: "thinkgos", key: "thinkgos"},
			want: "f9cebe2044ea375cff1a46f4dc05eb15ff9870ee",
		},
		{
			name: "sha224",
			args: args{method: MethodSha224, s: "thinkgos", key: "thinkgos"},
			want: "4c95a98768fabdff9756e2e92eda74ca062e00532c8c42eb67481701",
		},
		{
			name: "sha256",
			args: args{method: MethodSha256, s: "thinkgos", key: "thinkgos"},
			want: "e9403e3a615fad72d1dd1fe90c225cbec4ba81a03e5474d91a72844d2218954f",
		},
		{
			name: "sha384",
			args: args{method: MethodSha384, s: "thinkgos", key: "thinkgos"},
			want: "ba0c2ef006f64e7db43fb085abac27e960ed1c43e4604838c10ead6ffaa31dfd139f66343e1db84027271a267428ebbd",
		},
		{
			name: "sha512",
			args: args{method: MethodSha512, s: "thinkgos"},
			want: "c2181a3b42befba66ac95cf9fc6e11971c8bac0ec25bbf6805342b8166dd450e18ca5872e011ef1dd56bb960d96e7f93e1d2593d84f78e8e9a88892630393ce2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hmac(tt.args.method, tt.args.s, tt.args.key); got != tt.want {
				t.Errorf("Hmac(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
func TestAddSlashes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "add slash",
			args: args{s: "Is your name O'Reilly?"},
			want: `Is your name O\'Reilly?`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSlashes(tt.args.s); got != tt.want {
				t.Errorf("AddSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripSlashes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "strip slash",
			args: args{s: `Is your name O\'reilly?`},
			want: "Is your name O'reilly?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripSlashes(tt.args.s); got != tt.want {
				t.Errorf("StripSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuoteMeta(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "quote",
			args: args{s: "Hello world. (can you hear me?)"},
			want: `Hello world\. \(can you hear me\?\)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuoteMeta(tt.args.s); got != tt.want {
				t.Errorf("QuoteMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
