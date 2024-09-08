package patternTool

import (
	"fmt"
	"server-v2/card/cardSample"
	"testing"
)

func TestHasNSame(t *testing.T) {
	three := cardSample.TestHasThreeSame()
	four := cardSample.TestHasFourSame()

	fmt.Println("origin cards:", three)
	threeSameFlag, threeSames := HasNSame(three, 3)
	fmt.Println("has three same:", threeSameFlag)
	fmt.Println("three same cards:", threeSames)
	fmt.Println("=====================================")

	fmt.Println("origin cards:", four)
	fourSameFlag, fourSames := HasNSame(four, 4)
	fmt.Println("has four same:", fourSameFlag)
	fmt.Println("four same cards:", fourSames)
	fmt.Println("=====================================")

	fmt.Println("origin cards:", four)
	twoSameFlag, twoSames := HasNSame(four, 2)
	fmt.Println("has two same:", twoSameFlag)
	fmt.Println("four same cards:", twoSames)
}
