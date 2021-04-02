package number

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := Max(3, 44, 5, 6, 6, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001))
	fmt.Println(max)
}

func TestMin(t *testing.T) {
	max := Min(3, 44, 5, 6, 6, 2, 4, 9, 31, 23, 8, 3453, float64(4.999999999), float64(9.99999999), float64(2.0000001))
	fmt.Println(max)
}
