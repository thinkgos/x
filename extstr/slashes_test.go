package extstr

import (
	"testing"
)

func TestAddSlashes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			`add slash '`,
			"Is your name O'Reilly?",
			`Is your name O\'Reilly?`,
		},
		{
			`add slash "`,
			`Is your "name" Reilly?`,
			`Is your \"name\" Reilly?`,
		},
		{
			`add slash \`,
			`Is your name \Reilly?`,
			`Is your name \\Reilly?`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSlashes(tt.s); got != tt.want {
				t.Errorf("AddSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripSlashes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			`add slash '`,

			`Is your name O\'Reilly?`,
			"Is your name O'Reilly?",
		},
		{
			`add slash "`,

			`Is your \"name\" Reilly?`,
			`Is your "name" Reilly?`,
		},
		{
			`add slash \`,
			`Is your name \\Reilly?`,
			`Is your name \Reilly?`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripSlashes(tt.s); got != tt.want {
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
