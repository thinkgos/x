package extime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	ts := time.Now()
	t.Log(Microsecond(ts))
	t.Log(Millisecond(ts))
	t.Log(Time(Millisecond(ts)))
	Sleep(1)
	USleep(10)
	MSleep(10)
}

func TestValid(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常 ",
			args{
				2019,
				7,
				31,
			},
			true,
		},
		{
			"2月day不正常 ",
			args{
				2019,
				2,
				31,
			},
			false,
		},
		{
			"0月不正常 ",
			args{
				2019,
				0,
				31,
			},
			false,
		},
		{
			"13月不正常 ",
			args{
				2019,
				13,
				31,
			},
			false,
		},
		{
			"0年不正常 ",
			args{
				0,
				1,
				31,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDays(t *testing.T) {
	from := time.Date(2011, time.January, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2012, time.January, 4, 12, 1, 2, 100000, time.Local)
	t.Log(Days(to.Sub(from)))
}

func TestYearDays(t *testing.T) {
	require.Equal(t, 366, YearDays(2020))
	require.Equal(t, 365, YearDays(2019))
}

func TestMonthDays(t *testing.T) {
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1月", args{2020, time.January}, 31},
		{"闰年2月", args{2020, time.February}, 29},
		{"非闰年2月", args{2019, time.February}, 28},
		{"3月", args{2020, time.March}, 31},
		{"4月", args{2020, time.April}, 30},
		{"5月", args{2020, time.May}, 31},
		{"6月", args{2020, time.June}, 30},
		{"7月", args{2020, time.July}, 31},
		{"8月", args{2020, time.August}, 31},
		{"9月", args{2020, time.September}, 30},
		{"10月", args{2020, time.October}, 31},
		{"11月", args{2020, time.November}, 30},
		{"12月", args{2020, time.December}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthDays(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("MonthDays() = %v, want %v", got, tt.want)
			}
		})
	}
	require.Panics(t, func() {
		MonthDays(2020, 13)
	})
}

func TestMonthDays2(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1月", args{time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"2月", args{time.Date(2020, time.February, 1, 0, 0, 0, 0, time.Local)}, 29},
		{"3月", args{time.Date(2020, time.March, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"4月", args{time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"5月", args{time.Date(2020, time.May, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"6月", args{time.Date(2020, time.June, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"7月", args{time.Date(2020, time.July, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"8月", args{time.Date(2020, time.August, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"9月", args{time.Date(2020, time.September, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"10月", args{time.Date(2020, time.October, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"11月", args{time.Date(2020, time.November, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"12月", args{time.Date(2020, time.December, 1, 0, 0, 0, 0, time.Local)}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthDays2(tt.args.t); got != tt.want {
				t.Errorf("MonthDays2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate(t *testing.T) {
	type args struct {
		t      time.Time
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{
				time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
				"Y-m-d H:i:s",
			},
			"2012-11-22 21:28:10",
		},
		{
			"",
			args{
				time.Date(2012, 11, 22, 0, 0, 0, 0, time.Local),
				"Y-m-d",
			},
			"2012-11-22",
		},
		{
			"",
			args{
				time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
				"Y-m-d H:i:s",
			},
			"2012-11-22 21:28:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Date(tt.args.t, tt.args.layout); got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNow(t *testing.T) {
	t.Log(Now("r"))
	t.Log(Now("Y-m-d"))
	t.Log(Now("Y-m-d H:i:s"))
	t.Log(Now("Y-m-d H:i:s T"))
	t.Log(Now("Y-m-d H:i:s P"))
	t.Log(Now("Y-m-d H:i:s O"))
	t.Log(Now("Y-m-d H:i:s T O"))
	t.Log(Now("Y-m-d H:i:s T P"))
	t.Log(Now("Y-m-d H:i:s T P O"))
	t.Log(Now("Y-m-d H:i:s a T P O"))
	t.Log(Now("Y-m-d H:i:s A T P O"))
}

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  time.Time
	}{
		{
			"empty value",
			"",
			time.Time{},
		},
		{
			"2012-11-22 21:28:10 +0000 +0000",
			"2012-11-22 21:28:10 +0000 +0000",
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.UTC),
		},
		{
			"2012-11-22 +0000 +0000",
			"2012-11-22 +0000 +0000",
			time.Date(2012, 11, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			"2012-11-22 21:28:10 +0800 CST",
			"2012-11-22 21:28:10 +0800 CST",
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.FixedZone("CST", 28800)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.value); !got.Equal(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
	require.Panics(t, func() {
		Parse("20")
	})
}

func TestParseLocation(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  time.Time
	}{
		{
			"empty value",
			"",
			time.Time{}},
		{
			"2012-11-22 21:28:10",
			"2012-11-22 21:28:10",
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
		},
		{
			"2012-11-22",
			"2012-11-22",
			time.Date(2012, 11, 22, 0, 0, 0, 0, time.Local),
		},
		{
			"2012-11-22 21:28:10",
			"2012-11-22 21:28:10",
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLocation(tt.value); !got.Equal(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testDate = Parse("2020-03-10 23:04:35")

func TestStartOfDay(t *testing.T) {
	res := StartOfDay(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-10 00:00:00", str)
}

func TestEndOfDay(t *testing.T) {
	res := EndOfDay(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-10 23:59:59", str)
}

func TestStartOfMonth(t *testing.T) {
	res := StartOfMonth(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-01 00:00:00", str)
}

func TestEndOfMonth(t *testing.T) {
	res := EndOfMonth(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-31 23:59:59", str)
}

func TestStartOfYear(t *testing.T) {
	res := StartOfYear(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-01-01 00:00:00", str)
}

func TestEndOfYear(t *testing.T) {
	res := EndOfYear(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-12-31 23:59:59", str)
}

func TestStartOfWeek(t *testing.T) {
	res := StartOfWeek(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-09 00:00:00", str)

	d := Parse("2020-03-08 23:04:35")
	res = StartOfWeek(d, time.Tuesday)
	str = Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-03 00:00:00", str)
}

func TestEndOfWeek(t *testing.T) {
	res := EndOfWeek(testDate)
	str := Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-15 23:59:59", str)

	d := Parse("2020-03-08 23:04:35")
	res = EndOfWeek(d, time.Tuesday)
	str = Date(res, "Y-m-d H:i:s")
	require.Equal(t, "2020-03-09 23:59:59", str)
}
