package time

import (
	"fmt"
	"testing"
)

func TestFormatToUnix(t *testing.T) {
	fmt.Println(NewUnix(1617264318).ToFormat())
}
