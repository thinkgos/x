package habit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHideCard(t *testing.T) {
	require.Equal(t, "******", HideCard(""))
	require.Equal(t, "****", HideCard("1234"))
	require.Equal(t, "1234*", HideCard("12345"))
	require.Equal(t, "1234*****", HideCard("123456789"))
	require.Equal(t, "1234********345", HideCard("123456789012345"))
}

func BenchmarkHideCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HideCard("123456789")
	}
}

func TestHideMobile(t *testing.T) {
	require.Equal(t, "******", HideMobile(""))
	require.Equal(t, "*****", HideMobile("12345"))
	require.Equal(t, "123**6789", HideMobile("123456789"))
	require.Equal(t, "123********2345", HideMobile("123456789012345"))
}

func BenchmarkHideMobile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HideMobile("123456789")
	}
}

func TestHideName(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "", "******"},
		{"", "李", "李"},
		{"", "李四", "李*"},
		{"", "张三丰", "张**"},
		{"", "公孙先生", "公孙**"},
		{"", "helloWorld", "hel****rld"},
		{"", "北京搜狗科技公司", "北京****公司"},
		{"", "北京搜狗科技发展有限公司", "北京搜******限公司"},
		{"", "工商发展银行深圳南山科苑梅龙路支行", "工商发展*********龙路支行"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HideName(tt.s); got != tt.want {
				t.Errorf("HideName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHideName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HideName("北京搜狗科技发展有限公司")
	}
}
