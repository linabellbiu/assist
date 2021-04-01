package time

import (
	"fmt"
	"testing"
)

func TestFormatToUnix(t *testing.T) {
	fmt.Println(NewUnix(1617264318).ToFormat())
}

func TestNewFormat(t *testing.T) {
	tt, err := NewFormat("2021-04-01 12:00:00")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tt.ToIso8601())
}
