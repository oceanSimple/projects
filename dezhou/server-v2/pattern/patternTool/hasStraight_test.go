package patternTool

import (
	"fmt"
	"server-v2/card/cardSample"
	"testing"
)

func TestGetInitialArray(t *testing.T) {
	c := cardSample.TestHasStraight()
	fmt.Println("originalArray:", c)
	fmt.Println("initialArray:", getInitialArray(c))
}

func TestSpecialStraight(t *testing.T) {
	c := cardSample.TestHasSpecialStraight()
	fmt.Println("originalArray:", c)
	fmt.Println("specialStraight:", getSpecialStraight(c))
}

func TestHasStraight(t *testing.T) {
	c := cardSample.TestHasStraight()
	fmt.Println("originalArray:", c)
	straight, i := HasStraight(c)
	fmt.Println("has straight:", straight)
	fmt.Println("result:", i)

	fmt.Println("=====================================")

	c = cardSample.TestHasSpecialStraight()
	fmt.Println("originalArray:", c)
	straight, i = HasStraight(c)
	fmt.Println("has straight:", straight)
	fmt.Println("result:", i)
}
