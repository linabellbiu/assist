package number

import (
	"testing"
)

func TestCmp_Greater(t *testing.T) {
	res := Cpm.Greater(3, 2)
	t.Log(res)
}

func TestMax(t *testing.T) {
	max := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).Max()
	t.Log(max)
}

func TestMin(t *testing.T) {
	min := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).Min()
	t.Log(min)
}

func TestNumber_OrderAsc(t *testing.T) {
	asc := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).OrderAsc()
	t.Log(asc)
}

func TestNumber_OrderDesc(t *testing.T) {
	desc := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).OrderDesc()
	t.Log(desc)
}
