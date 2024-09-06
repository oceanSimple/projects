package card

type ColorEnum struct {
	Hearts   ColorObj
	Diamonds ColorObj
	Clubs    ColorObj
	Spades   ColorObj
}

type ColorObj struct {
	Text    string
	Picture string
}

func getColor() ColorEnum {
	return ColorEnum{
		Hearts:   ColorObj{"Hearts", "♥"},
		Diamonds: ColorObj{"Diamonds", "♦"},
		Clubs:    ColorObj{"Clubs", "♣"},
		Spades:   ColorObj{"Spades", "♠"},
	}
}

func getColors() []ColorObj {
	colors := getColor()
	return []ColorObj{
		colors.Hearts,
		colors.Diamonds,
		colors.Clubs,
		colors.Spades,
	}
}
