package card

type NumberObj struct {
	// 文字描述：A、2、3、4、5、6、7、8、9、10、J、Q、K
	Text string
	// 权重：A=14、2=2、3=3、4=4、5=5、6=6、7=7、8=8、9=9、10=10、J=11、Q=12、K=13
	Weight int
	// 特殊权重，A=1，用来匹配顺子：A、2、3、4、5
	ExtraWeight int
}

// 数字枚举，用来方便地获取数字对象
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

func getNumberEnum() NumberEnum {
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

func getNumberArray() []NumberObj {
	numbers := getNumberEnum()
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
