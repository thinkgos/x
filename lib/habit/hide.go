package habit

import (
	"strings"
)

const sixStar = "******"

// HideCard 隐藏证件号码.
func HideCard(card string) string {
	length := len(card)
	switch {
	case length == 0:
		return sixStar
	case length > 4 && length <= 10:
		return card[:4] + strings.Repeat("*", length-4)
	case length > 10:
		return card[:4] + strings.Repeat("*", length-7) + card[(length-3):]
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
		return mobile[:3] + strings.Repeat("*", length-7) + mobile[length-4:]
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
		return string(runs[:1]) + strings.Repeat("*", length-1)
	case length < 5:
		return string(runs[:2]) + strings.Repeat("*", length-2)
	case length < 10:
		return string(runs[:2]) + strings.Repeat("*", length-4) + string(runs[length-2:])
	case length < 16:
		return string(runs[:3]) + strings.Repeat("*", length-6) + string(runs[length-3:])
	default:
		return string(runs[:4]) + strings.Repeat("*", length-8) + string(runs[length-4:])
	}
}
