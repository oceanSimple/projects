package pattern

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"server-v1/static/cardSample"
	"testing"
)

// 皇家同花顺：♠A-♥A-♠K-♥K-♠Q-♠J-♠10
func TestRoyalFlushStraight(t *testing.T) {
	c := cardSample.RoyalFlushStraight()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 8, weight)
}

// 皇家同花顺：♠A-♠K-♥K-♥Q-♠Q-♠J-♠10
func TestRoyalFlushStraight1(t *testing.T) {
	c := cardSample.RoyalFlushStraight1()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 8, weight)
}

// 最小同花顺：♠A-♥A-♠5-♥4-♠4-♠3-♠2
func TestStraightFlushMin(t *testing.T) {
	c := cardSample.StraightFlushMin()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 8, weight)
}

// 同花顺：♠9-♠8-♠7-♠6-♠5-♠4-♠3
func TestStraightFlush(t *testing.T) {
	c := cardSample.StraightFlush()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 8, weight)
}

// 四条： ♠A-♥A-♥K-♠K-♦K-♣K-♠Q
func TestFourOfAKind(t *testing.T) {
	c := cardSample.FourOfAKind()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 7, weight)
}

// 葫芦：♠A-♥A-♦A-♥K-♠K-♦K-♣Q
func TestFullHouse(t *testing.T) {
	c := cardSample.FullHouse()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 6, weight)
}

// 非顺子的同花：♠K-♠J-♥10-♠8-♦5-♠4-♠3
func TestFlush(t *testing.T) {
	c := cardSample.Flush()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 5, weight)
}

// 非同花的顺子：♠A-♥K-♠Q-♦J-♥10-♣9-♠8
func TestStraight(t *testing.T) {
	c := cardSample.Straight()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 4, weight)
}

// 三条：♠A-♥A-♦A-♠K-♣Q-♠J-♠7
func TestThreeOfAKind(t *testing.T) {
	c := cardSample.ThreeOfAKind()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 3, weight)
}

// 两对：♠A-♥A-♦K-♠K-♣Q-♠J-♠7
func TestTwoPairs(t *testing.T) {
	c := cardSample.TwoPairs()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 2, weight)
}

// 一对：♠A-♥A-♦K-♠Q-♥J-♠4-♠3
func TestOnePair(t *testing.T) {
	c := cardSample.OnePair()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 1, weight)
}

// 散排：♠A-♥Q-♦J-♠8-♦6-♠3-♠2
func TestHighCard(t *testing.T) {
	c := cardSample.HighCard()
	fmt.Println("origin cards:", c)
	weight, arr := Judge(c)
	fmt.Println("weight:", weight, "arr:", arr)
	assert.Equal(t, 0, weight)
}
