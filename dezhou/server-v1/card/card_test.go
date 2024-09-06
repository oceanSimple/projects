package card

import (
	"fmt"
	"testing"
)

func TestGetCardByString(t *testing.T) {
	c1 := GetCardByString("1-1")
	c2 := GetCardByString("2-2")
	c3 := GetCardByString("3-3")
	c4 := GetCardByString("4-11")

	fmt.Println(c1, c2, c3, c4)
}
