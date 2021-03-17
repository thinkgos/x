package extime

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Microsecond time.Time 转为 微秒
func Microsecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

// Millisecond time.Time 转为 毫秒
func Millisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// NowUS time.Now() 转为 微秒
func NowUS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NowMS time.Now() 转为 毫秒
func NowMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
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

// Valid 检查是否正常的日期.
func Valid(year, month, day int) bool {
	return month >= 1 && month <= 12 &&
		day >= 1 && day <= 31 &&
		year >= 1 && year <= math.MaxInt32 &&
		day <= MonthDays(year, time.Month(month))
}

// Days time.Duration转化为天数
func Days(d time.Duration) float64 {
	day := d / (24 * time.Hour)
	nsec := d % (24 * time.Hour)
	return float64(day) + float64(nsec)/(24*60*60*1e9)
}

// IsLeapYear 是否闰年
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

// MonthDays 所在年份月份的天数
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

// MonthDays2 t 所在时间月份的天数
func MonthDays2(t time.Time) int { return MonthDays(t.Year(), t.Month()) }

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
func Now(layout string) string { return Date(time.Now(), layout) }

// Parse parse value use PHP time format.
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

// ParseLocation parse location
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

// StartOfDay 获取日期中当天的开始时间.
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 获取日期中当天的结束时间.
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// StartOfMonth 获取日期中当月的开始时间.
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 获取日期中当月的结束时间.
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// StartOfYear 获取日期中当年的开始时间.
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 获取日期中当年的结束时间.
func EndOfYear(t time.Time) time.Time {
	return StartOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// StartOfWeek 获取日期中当周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func StartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	weekStart := time.Monday
	if len(weekStartDay) > 0 {
		weekStart = weekStartDay[0]
	}

	// 当前是周几
	weekday := int(date.Weekday())
	if weekStart != time.Sunday {
		weekStartDayInt := int(weekStart)

		if weekday < weekStartDayInt {
			weekday += 7 - weekStartDayInt
		} else {
			weekday -= weekStartDayInt
		}
	}

	return time.Date(date.Year(), date.Month(), date.Day()-weekday, 0, 0, 0, 0, date.Location())
}

// EndOfWeek 获取日期中当周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func EndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return StartOfWeek(date, weekStartDay...).AddDate(0, 0, 7).Add(-time.Nanosecond)
}
