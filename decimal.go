package assist

import (
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

func StringToFloat64(value string) (f float64) {
	f, _ = strconv.ParseFloat(value, 64)
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

func Float64ToInt64(value float64) (i int64) {
	i = decimal.NewFromFloat(value).IntPart()
	return
}

func Float64ToString(value float64) (s string) {
	return strconv.FormatFloat(value, 'f', -1, 64) //float64
}

func RoundFloat64(value float64, exp int) (num float64) {
	//floatStr := fmt.Sprintf("%."+strconv.Itoa(exp)+"f", value)
	//num, _ = strconv.ParseFloat(floatStr, 64)
	//return
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

var A Accuracy = func() float64 { return 0.000000001 }

type Accuracy func() float64

//=
func (this Accuracy) Equal(a, b float64) bool {
	return math.Abs(a-b) < this()
}

//>
func (this Accuracy) Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > this()
}

//<
func (this Accuracy) Smaller(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > this()
}

//>=
func (this Accuracy) GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < this()
}

//<=
func (this Accuracy) SmallerOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < this()
}
