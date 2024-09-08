package card

import "strings"

var (
	NumberConst = getNumberEnum() // 数字对象枚举
	ColorConst  = getColorEnum()  // 颜色对象枚举
)

// 根据字符串获取一张牌
// 例如：♠A
func GetCardByString(str string) Card {
	var color ColorObj
	var number NumberObj

	splits := strings.Split(str, "")

	switch splits[0] {
	case "♠":
		color = ColorConst.Spades
	case "♥":
		color = ColorConst.Hearts
	case "♣":
		color = ColorConst.Clubs
	case "♦":
		color = ColorConst.Diamonds
	}
	switch splits[1] {
	case "A":
		number = NumberConst.Ace
	case "2":
		number = NumberConst.Two
	case "3":
		number = NumberConst.Three
	case "4":
		number = NumberConst.Four
	case "5":
		number = NumberConst.Five
	case "6":
		number = NumberConst.Six
	case "7":
		number = NumberConst.Seven
	case "8":
		number = NumberConst.Eight
	case "9":
		number = NumberConst.Nine
	case "1": // 特殊处理，因为A用A对应，因此可以用1对应10而不引起歧义
		number = NumberConst.Ten
	case "J":
		number = NumberConst.Jack
	case "Q":
		number = NumberConst.Queen
	case "K":
		number = NumberConst.King
	}
	return Card{Color: color, Number: number}
}
