package pattern

import (
	"fmt"
	"server-v2/card/cardSample"
	"testing"
)

func TestStraightFlush_Judge(t *testing.T) {
	judger := NewStraightFlush()

	c := cardSample.RoyalStraightFlush()
	fmt.Println("origin cards:", c)
	judge, cards := judger.Judge(c)
	fmt.Println("judge:", judge)
	fmt.Println("cards:", cards)

	fmt.Println("=====================================")

	c = cardSample.StraightFlush()
	fmt.Println("origin cards:", c)
	judge, cards = judger.Judge(c)
	fmt.Println("judge:", judge)
	fmt.Println("cards:", cards)
}
