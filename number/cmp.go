package number

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"sort"
)

type Cmp func(float64, float64) int

type number struct {
	Cmp Cmp
	n   []interface{}
}

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

func (n number) Len() int { return len(n.n) }

func (n number) Less(i, j int) bool { return n.Cmp.GreaterOrEqual(toFloat(n.n[i]), toFloat(n.n[j])) }

func (n number) Swap(i, j int) { n.n[i], n.n[j] = n.n[j], n.n[i] }

//从大到小
func OrderDesc() {

}

//从小到大
func OrderAsc() {

}

//找出数组中最大的数字
//数组可以包含float64,float32,int,int64,int32
func Max(n ...interface{}) float64 {
	num := number{
		Cpm,
		n,
	}
	sort.Sort(num)
	return toFloat(num.n[0])
}

//找出数组中最小的数字
//数组可以包含float64,float32,int,int64,int32
func Min(n ...interface{}) float64 {
	num := number{
		Cpm,
		n,
	}
	sort.Sort(num)
	return toFloat(num.n[len(n)-1])
}

func toFloat(val interface{}) float64 {
	switch val.(type) {
	case float64:
		return val.(float64)
	case float32:
		return float64(val.(float32))
	case string:
		return StringToFloat64(val.(string))
	case int64:
		return Int64ToFloat64(val.(int64))
	case int:
		return IntToFloat64(val.(int))
	case int32:
		return Int64ToFloat64(int64(val.(int32)))
	default:
		panic(fmt.Sprintf("%s to string:Unsupported type conversion", reflect.TypeOf(val)))
	}
	return 0
}
