package card

type ColorObj struct {
	// 文字描述：Hearts, Diamonds, Clubs, Spades
	Text string
	// 符号描述：♥, ♦, ♣, ♠
	Picture string
}

// 颜色枚举
type ColorEnum struct {
	Hearts   ColorObj
	Diamonds ColorObj
	Clubs    ColorObj
	Spades   ColorObj
}

func getColorEnum() ColorEnum {
	return ColorEnum{
		Hearts:   ColorObj{"Hearts", "♥"},
		Diamonds: ColorObj{"Diamonds", "♦"},
		Clubs:    ColorObj{"Clubs", "♣"},
		Spades:   ColorObj{"Spades", "♠"},
	}
}

func getColorArray() []ColorObj {
	colors := getColorEnum()
	return []ColorObj{
		colors.Hearts,
		colors.Diamonds,
		colors.Clubs,
		colors.Spades,
	}
}
