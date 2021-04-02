package number

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := NewNumbers(3, 44, 5, 6, 6, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001)).Max()
	fmt.Println(max)
}

func TestMin(t *testing.T) {
	max := NewNumbers(3, 44, 5, 6, 6, 2, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001)).Min()
	fmt.Println(max)
}

func TestNumber_OrderAsc(t *testing.T) {
	asc := NewNumbers(3, 44, 5, 6, 6, 2, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001)).OrderAsc()
	fmt.Println(asc)
}

func TestNumber_OrderDesc(t *testing.T) {
	desc := NewNumbers(3, 44, 5, 6, 6, 2, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001)).OrderDesc()
	fmt.Println(desc)
}
