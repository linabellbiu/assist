package number

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func StringToFloat64(value string) (f float64) {
	if value == "" {
		value = "0"
	}
	f, _ = strconv.ParseFloat(value, 64)
	return
}

func Int64ToFloat64(value int64) (f float64) {
	d := decimal.NewFromInt(value)
	f, _ = d.Float64()
	return
}

func StringToInt64(value string) (i int64) {
	if value == "" {
		value = "0"
	}
	d, _ := decimal.NewFromString(value)
	i = d.IntPart()
	return
}

func StringToInt(value string) (i int) {
	if value == "" {
		value = "0"
	}
	i, _ = strconv.Atoi(value)
	return
}

func Float64ToInt64(value float64) (i int64) {
	i = decimal.NewFromFloat(value).IntPart()
	return
}

func Float64ToString(value float64) (s string) {
	return strconv.FormatFloat(value, 'f', -1, 64) //float64
}

func RoundFloat64(value float64, exp int) (num float64) {
	num, _ = decimal.NewFromFloat(value).Round(int32(exp)).Float64()
	return
}

func Int64ToString(value int64) (s string) {
	s = decimal.New(value, 0).String()
	return
}

func IntToFloat64(value int) (f float64) {
	s, _ := decimal.NewFromString(strconv.Itoa(value))
	f = StringToFloat64(s.String())
	return
}
