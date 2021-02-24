package habit

import (
	"strings"
)

const sixStar = "******"

// HideCard 隐藏证件号码.证件号码都为数字+字母
func HideCard(card string) string {
	length := len(card)
	switch {
	case length == 0:
		return sixStar
	case length > 4 && length <= 10:
		return BuildHideString(card[:4], "", length-4)
	case length > 10:
		return BuildHideString(card[:4], card[length-3:], length-7)
	default:
		return strings.Repeat("*", length)
	}
}

// HideMobile 隐藏手机号.
func HideMobile(mobile string) string {
	length := len(mobile)
	switch {
	case length == 0:
		return sixStar
	case length > 7:
		return BuildHideString(mobile[:3], mobile[length-4:], length-7)
	default:
		return strings.Repeat("*", length)
	}
}

// HideName 隐藏真实名称(如姓名、账号、公司等).
func HideName(s string) string {
	if s == "" {
		return sixStar
	}
	runs := []rune(s)
	length := len(runs)
	switch {
	case length <= 3:
		return BuildHideString(string(runs[:1]), "", length-1)
	case length < 5:
		return BuildHideString(string(runs[:2]), "", length-2)
	case length < 10:
		return BuildHideString(string(runs[:2]), string(runs[length-2:]), length-4)
	case length < 16:
		return BuildHideString(string(runs[:3]), string(runs[length-3:]), length-6)
	default:
		return BuildHideString(string(runs[:4]), string(runs[length-4:]), length-8)
	}
}

// BuildHideString 生成隐藏字符串,中间使用 '*' 表示
func BuildHideString(prefix, suffix string, midStarRepeatCnt int) string {
	var b strings.Builder

	b.Grow(len(prefix) + len(suffix) + midStarRepeatCnt)
	b.WriteString(prefix)
	for b.Len() < len(prefix)+midStarRepeatCnt {
		b.WriteString("*")
	}
	b.WriteString(suffix)
	return b.String()
}
