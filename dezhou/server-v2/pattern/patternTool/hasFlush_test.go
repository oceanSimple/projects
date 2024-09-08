package patternTool

import (
	"fmt"
	"server-v2/card/cardSample"
	"testing"
)

func TestHasFlush(t *testing.T) {
	cards := cardSample.TestHasFlush()
	hasFlush, flushes := HasFlush(cards)
	fmt.Println("origin cards:", cards)
	fmt.Println("has flush:", hasFlush)
	fmt.Println("flushes:", flushes)
}
