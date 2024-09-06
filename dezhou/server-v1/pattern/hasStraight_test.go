package pattern

import (
	"fmt"
	"server-v1/static/cardSample"
	"testing"
)

func TestHasStraight(t *testing.T) {
	c := cardSample.RoyalFlushStraight()
	fmt.Println("origin cards:", c)
	_, result := HasStraight(c)
	fmt.Println("result:")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

	fmt.Println("=====================================")
	c1 := cardSample.RoyalFlushStraight1()
	fmt.Println("origin cards:", c1)
	_, result1 := HasStraight(c1)
	fmt.Println("result:")
	for i := 0; i < len(result1); i++ {
		fmt.Println(result1[i])
	}

	fmt.Println("=====================================")
	c2 := cardSample.StraightFlushMin()
	fmt.Println("origin cards:", c2)
	_, result2 := HasStraight(c2)
	fmt.Println("result:")
	for i := 0; i < len(result2); i++ {
		fmt.Println(result2[i])
	}
}

// 非同花的顺子：♠A-♥K-♠Q-♦J-♥10-♣9-♠8
func TestHasStraight2(t *testing.T) {
	c := cardSample.Straight()
	fmt.Println("origin cards:", c)
	f, result := HasStraight(c)
	fmt.Println("result:", f)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
