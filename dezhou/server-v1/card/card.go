package card

import (
	"fmt"
	"server-v1/output"
	"strings"
)

type Card struct {
	Color  ColorObj
	Number NumberObj
}

func (c *Card) String() string {
	return fmt.Sprintf("{%s, %s}", c.Color.Picture, c.Number.Text)
}

func (c *Card) DyeString() string {
	if c.Color.Picture == "♠" || c.Color.Picture == "♣" {
		return fmt.Sprintf("{%s, %s}", output.DyeText(c.Color.Picture, output.Grey), c.Number.Text)
	} else {
		return fmt.Sprintf("{%s, %s}", output.DyeText(c.Color.Picture, output.Red), c.Number.Text)
	}
}

// pattern: n-n
// the first n is the color of the card: [1-4] ♠♥♣♦
// the second n is the number of the card: [1-13] A-2-3-4-5-6-7-8-9-10-J-Q-K
func GetCardByString(str string) *Card {
	var color ColorObj
	var number NumberObj

	splits := strings.Split(str, "-")

	switch splits[0] {
	case "1":
		color = Color.Spades
	case "2":
		color = Color.Hearts
	case "3":
		color = Color.Clubs
	case "4":
		color = Color.Diamonds
	}
	switch splits[1] {
	case "1":
		number = Number.Ace
	case "2":
		number = Number.Two
	case "3":
		number = Number.Three
	case "4":
		number = Number.Four
	case "5":
		number = Number.Five
	case "6":
		number = Number.Six
	case "7":
		number = Number.Seven
	case "8":
		number = Number.Eight
	case "9":
		number = Number.Nine
	case "10":
		number = Number.Ten
	case "11":
		number = Number.Jack
	case "12":
		number = Number.Queen
	case "13":
		number = Number.King
	}
	return &Card{Color: color, Number: number}
}
