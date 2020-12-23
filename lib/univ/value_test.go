package univ

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValues_Encode(t *testing.T) {
	v := make(Values)
	v.Add("aa", "a1")
	v.Add("aa", "a2")
	v.Add("bb", "bb")
	v.Add("cc", "cc")
	require.Equal(t, "aa=a1&aa=a2&bb=bb&cc=cc", v.Encode("=", "&"))

	// 不存在
	require.Equal(t, "", v.Get("dd"))
	// 存在,取处一个值
	require.Equal(t, "a1", v.Get("aa"))
	// 删除
	v.Del("aa")
	require.Equal(t, "", v.Get("aa"))
	// 设置覆盖
	v.Set("aa", "aa")
	require.Equal(t, "aa=aa&bb=bb&cc=cc", v.Encode("=", "&"))

	var v1 Values

	require.Equal(t, "", v1.Get("aa"))
	require.Equal(t, "", v1.Encode("=", "&"))
}

func TestParseValue(t *testing.T) {
	vs := ParseValue("aa=aa&bb=bb&cc=cc&&=dd", "=", "&")
	want := map[string]string{
		"aa": "aa",
		"bb": "bb",
		"cc": "cc",
	}

	for k, v := range want {
		require.Equal(t, v, vs.Get(k))
	}
	t.Log(vs.Encode("=", "&"))
}
