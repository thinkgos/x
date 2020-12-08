package exttime

import (
	"fmt"
	"strings"
	"time"
)

// Microsecond time.Time 转为 微秒
func Microsecond(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Microsecond)
}

// Millisecond time.Time 转为 毫秒
func Millisecond(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Millisecond)
}

// Time 毫秒转time.Time
func Time(msec int64) time.Time {
	return time.Unix(msec/1000, (msec%1000)*int64(time.Millisecond))
}

// Sleep pauses the current goroutine for at least the second d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// MSleep pauses the current goroutine for at least the millisecond d.
// A negative or zero duration causes Sleep to return immediately.
func MSleep(t int64) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

// USleep pauses the current goroutine for at least the microsecond d.
// A negative or zero duration causes Sleep to return immediately.
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

// IsLeap 是否闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// YearDays 所在年份总天数
func YearDays(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// MonthDays 所在年份月份的总天数
func MonthDays(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July,
		time.August, time.October, time.December:
		return 31
	case time.February:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case time.April, time.June, time.September, time.November:
		return 30
	default:
		panic(fmt.Errorf("invalid month %v", month))
	}
}

// MonthDays2 t 所在时间月份的总天数
func MonthDays2(t time.Time) int {
	return MonthDays(t.Year(), t.Month())
}

// MillisecondsBetween 获取两个日期的间隔毫秒数.
func MillisecondsBetween(from, to time.Time) int64 {
	return int64(to.Sub(from) / (time.Millisecond))
}

// SecondsBetween 获取两个日期的间隔秒数.
func SecondsBetween(from, to time.Time) int64 {
	return int64(to.Sub(from) / (time.Second))
}

// MinutesBetween 获取两个日期的间隔分钟数.
func MinutesBetween(from, to time.Time) int64 {
	return int64(to.Sub(from) / (time.Minute))
}

// HoursBetween 获取两个日期的间隔小时数.
func HoursBetween(from, to time.Time) int64 {
	return int64(to.Sub(from) / (time.Hour))
}

// DaysBetween 获取两个日期的间隔天数.
func DaysBetween(from, to time.Time) int64 {
	return int64(to.Sub(from) / (24 * time.Hour))
}

// Date Format pattern rules.
var replacer = strings.NewReplacer([]string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", // A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}...)

// Date 跟 PHP 中 date 类似的使用方式
// layout 格式,如"Y-m-d H:i:s".
func Date(t time.Time, layout string) string {
	layout = replacer.Replace(layout)
	return t.Format(layout)
}

// Now 跟 PHP 中 date 类似的使用方式
// layout 格式,如"Y-m-d H:i:s".
func Now(layout string) string {
	return Date(time.Now(), layout)
}

func Parse(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}
	panic(err)
}

func ParseLocation(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return Parse(value)
}
