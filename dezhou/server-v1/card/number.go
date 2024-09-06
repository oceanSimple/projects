package card

type NumberEnum struct {
	Ace   NumberObj
	Two   NumberObj
	Three NumberObj
	Four  NumberObj
	Five  NumberObj
	Six   NumberObj
	Seven NumberObj
	Eight NumberObj
	Nine  NumberObj
	Ten   NumberObj
	Jack  NumberObj
	Queen NumberObj
	King  NumberObj
}

type NumberObj struct {
	Text        string
	Weight      int
	ExtraWeight int //another weight for special case：A
}

func getNumber() NumberEnum {
	return NumberEnum{
		Ace:   NumberObj{"A", 14, 1},
		Two:   NumberObj{"2", 2, 2},
		Three: NumberObj{"3", 3, 3},
		Four:  NumberObj{"4", 4, 4},
		Five:  NumberObj{"5", 5, 5},
		Six:   NumberObj{"6", 6, 6},
		Seven: NumberObj{"7", 7, 7},
		Eight: NumberObj{"8", 8, 8},
		Nine:  NumberObj{"9", 9, 9},
		Ten:   NumberObj{"10", 10, 10},
		Jack:  NumberObj{"J", 11, 11},
		Queen: NumberObj{"Q", 12, 12},
		King:  NumberObj{"K", 13, 13},
	}
}

func getNumbers() []NumberObj {
	numbers := getNumber()
	return []NumberObj{
		numbers.Ace,
		numbers.Two,
		numbers.Three,
		numbers.Four,
		numbers.Five,
		numbers.Six,
		numbers.Seven,
		numbers.Eight,
		numbers.Nine,
		numbers.Ten,
		numbers.Jack,
		numbers.Queen,
		numbers.King,
	}
}
