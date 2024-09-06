package pattern

import (
	"fmt"
	"server-v1/card"
	"server-v1/static/cardSample"
	"testing"
)

func TestHasFlush(t *testing.T) {
	flush := cardSample.A_J_10_5_4_3_2_NoFlush()
	fmt.Println(flush)
	done, result := HasFlush(flush)
	t.Log("A, J, 10, 5, 4, 3, 2 is flush:", done, result)
}

func TestHasFourSame(t *testing.T) {
	c1 := cardSample.K_Q_J_J_J_J_5()
	fmt.Println("origin cards:", c1)
	done, result := HasFourSame(c1)
	fmt.Println("result:", done, result)
}

func TestHasThreeSame(t *testing.T) {
	c1 := cardSample.K_Q_J_J_J_J_5()
	fmt.Println("origin cards:", c1)
	done, result := HasThreeSame(c1)
	fmt.Println("result:", done, result)
}

func TestHasTwoSame(t *testing.T) {
	c1 := cardSample.K_Q_J_J_J_J_5()
	fmt.Println("origin cards:", c1)
	done, result := HasTwoSame(c1)
	fmt.Println("result:", done, result)
}

func TestDeduplication(t *testing.T) {
	var c = [][]*card.Card{
		{&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts}, &card.Card{Number: card.Number.Ace, Color: card.Color.Clubs}},
		{&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts}, &card.Card{Number: card.Number.Ace, Color: card.Color.Spades}},
		{&card.Card{Number: card.Number.Eight, Color: card.Color.Hearts}, &card.Card{Number: card.Number.Eight, Color: card.Color.Hearts}},
	}
	fmt.Println("origin:", c)
	c = Deduplication(c)
	fmt.Println("result:", c)
}
