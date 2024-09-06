package card

import (
	"fmt"
	"math/rand"
	"time"
)

type CardBox struct {
	Box []*Card
}

// Create a new card box and shuffle it.
func NewCardBox() *CardBox {
	var box = &CardBox{getNewCardBox()}
	box.ShuffleCardBox()
	return box
}

// Shuffle the card box.
func (cardBox *CardBox) ShuffleCardBox() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(cardBox.Box), func(i, j int) {
		cardBox.Box[i], cardBox.Box[j] = cardBox.Box[j], cardBox.Box[i]
	})
}

// Get one card from the card box.
func (cardBox *CardBox) GetOneCard() *Card {
	card := cardBox.Box[0]
	cardBox.Box = cardBox.Box[1:]
	return card
}

// Print the card box in 4 rows, each row has 13 cards.
func (cardBox *CardBox) printCardBox() {
	for i := 0; i < len(cardBox.Box); i++ {
		fmt.Printf("%s ", cardBox.Box[i].String())
		if (i+1)%13 == 0 {
			fmt.Println()
		}
	}
}

func getNewCardBox() []*Card {
	var cardBox []*Card

	for _, color := range ColorSlice {
		for _, number := range NumberSlice {
			cardBox = append(cardBox, &Card{color, number})
		}
	}

	return cardBox
}
