package tool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"server-v2/card/cardSample"
	"testing"
)

func TestSortCardsByNumberWeight(t *testing.T) {
	c := cardSample.TestSort()
	fmt.Println("Before sorting:", PrintCards(c))
	SortCardsByNumberWeight(c)
	fmt.Println("After sorting:", PrintCards(c))
}

func TestSortCardsByNumberExtraWeight(t *testing.T) {
	c := cardSample.TestSort()
	fmt.Println("Before sorting:", PrintCards(c))
	SortCardsByNumberExtraWeight(c)
	fmt.Println("After sorting:", PrintCards(c))
}

func TestSortCardsByColorText(t *testing.T) {
	c := cardSample.TestSort()
	fmt.Println("Before sorting:", PrintCards(c))
	SortCardsByColorText(c)
	fmt.Println("After sorting:", PrintCards(c))
}

func TestEqualCards(t *testing.T) {
	c1 := cardSample.TestSort()
	c2 := cardSample.TestSort()
	fmt.Println("Equal:", EqualCards(c1, c2))
	assert.Equal(t, true, EqualCards(c1, c2))
}
