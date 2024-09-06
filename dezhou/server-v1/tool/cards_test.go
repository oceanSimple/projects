package tool

import (
	"server-v1/card"
	"server-v1/static/cardSample"
	"testing"
)

func TestSortCards(t *testing.T) {
	// 5, A, 4, K, J
	cards := []*card.Card{
		{
			Color:  card.Color.Hearts,
			Number: card.Number.Five,
		},
		{
			Color:  card.Color.Hearts,
			Number: card.Number.Ace,
		},
		{
			Color:  card.Color.Hearts,
			Number: card.Number.Four,
		},
		{
			Color:  card.Color.Hearts,
			Number: card.Number.King,
		},
		{
			Color:  card.Color.Hearts,
			Number: card.Number.Jack,
		},
	}
	t.Log("before sort: ", cards)
	SortCardsByNumber(cards)
	t.Log("after sort: ", cards)
}

func TestEqualCardSlice(t *testing.T) {
	c1 := cardSample.A_J_10_5_4_3_2()
	c2 := cardSample.A_J_10_5_4_3_2()
	t.Log("c1 == c2: ", EqualCardSlice(c1, c2))
}
