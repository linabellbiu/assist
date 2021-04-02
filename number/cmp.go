package number

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"sort"
)

type Cmp func(float64, float64) int

type number struct {
	Cmp   Cmp
	n     []interface{}
	order string
}

const orderDesc = "desc"
const orderAsc = "asc"

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

func (n number) Swap(i, j int) { n.n[i], n.n[j] = n.n[j], n.n[i] }

func (n number) Less(i, j int) bool {
	l, r := toFloat(n.n[i]), toFloat(n.n[j])
	if n.order == orderAsc {
		return n.Cmp.SmallerOrEqual(l, r)
	}
	return n.Cmp.GreaterOrEqual(l, r)
}

func NewNumbers(n ...interface{}) number {
	return number{
		Cmp: Cpm,
		n:   n,
	}
}

//从大到小排序
//数组可以包含float64,float32,int,int64,int32
func (n number) OrderDesc() []interface{} {
	n.order = orderDesc
	sort.Sort(n)
	return n.n
}

//从小到大排序
//数组可以包含float64,float32,int,int64,int32
func (n number) OrderAsc() []interface{} {
	n.order = orderAsc
	sort.Sort(n)
	return n.n
}

//找出数组中最大的数字
//数组可以包含float64,float32,int,int64,int32
func (n number) Max() float64 {
	_arr := n.OrderDesc()
	return toFloat(_arr[0])
}

//找出数组中最小的数字
//数组可以包含float64,float32,int,int64,int32
func (n number) Min() float64 {
	_arr := n.OrderAsc()
	return toFloat(_arr[0])
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
}
