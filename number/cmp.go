package number

import (
	"github.com/shopspring/decimal"
)

type Cmp func(float64, float64) int

var Cpm Cmp = func(l, r float64) int {
	_l := decimal.NewFromFloat(l)
	_r := decimal.NewFromFloat(r)
	return _l.Cmp(_r)
}

//=
func (c Cmp) Equal(l, r float64) bool {
	if c(l, r) == 0 {
		return true
	}
	return false
}

//>
func (c Cmp) Greater(l, r float64) bool {
	if c(l, r) == 1 {
		return true
	}
	return false
}

//<
func (c Cmp) Smaller(l, r float64) bool {
	if c(l, r) == -1 {
		return true
	}
	return false
}

//>=
func (c Cmp) GreaterOrEqual(l, r float64) bool {
	if res := c(l, r); res == 0 || res == 1 {
		return true
	}
	return false
}

//<=
func (c Cmp) SmallerOrEqual(l, r float64) bool {
	if res := c(l, r); res == 0 || res == -1 {
		return true
	}
	return false
}

func toFloat(n interface{}) float64 {
	switch n.(type) {
	case float64:
		return n.(float64)
	case string:
		return StringToFloat64(n.(string))
	case int64:
		return Int64ToFloat64(n.(int64))
	case int:
		return IntToFloat64(n.(int))
	default:

	}
	return 0
}
