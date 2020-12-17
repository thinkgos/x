package habit

import (
	"github.com/thinkgos/go-core-package/extstr"
)

// ParseIdsGroup 解析以','分隔的批量id
func ParseIdsGroup(s string) []int64 {
	return extstr.Split(s, ",")
}

// ParseIdsGroupInt 解析以','分隔的批量id
func ParseIdsGroupInt(s string) []int {
	return extstr.SplitInt(s, ",")
}

// IdsGroup 以','分隔的批量id为字符串
func IdsGroup(ids []int64) string {
	return extstr.Join(ids, ",")
}

// IdsGroupInt 以','分隔的批量id为字符串
func IdsGroupInt(ids []int) string {
	return extstr.JoinInt(ids, ",")
}
