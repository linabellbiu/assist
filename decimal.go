package assist

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func StringToFloat64(value string) (f float64) {
	d, _ := decimal.NewFromString(value)
	f, _ = d.Float64()
	return
}

func Int64ToFloat64(value int64) (f float64) {
	d := decimal.NewFromInt(value)
	f, _ = d.Float64()
	return
}

func StringToInt64(value string) (i int64) {
	d, _ := decimal.NewFromString(value)
	i = d.IntPart()
	return
}

func StringToInt(value string) (i int) {
	i, _ = strconv.Atoi(value)
	return
}

//测试
