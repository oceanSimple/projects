package card

import (
	"fmt"
	"math/rand"
	"time"
)

type CardBox struct {
	Box []Card // 牌盒
}

// 获得一副崭新的、未洗牌的扑克牌
func NewCardBox() *CardBox {
	return &CardBox{Box: getNewCardBox()}
}

// 获取牌盒顶部的一张牌
func (cardBox *CardBox) GetOneCard() Card {
	card := cardBox.Box[0]
	cardBox.Box = cardBox.Box[1:]
	return card
}

// 洗牌
func (cardBox *CardBox) ShuffleCardBox() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(cardBox.Box), func(i, j int) {
		cardBox.Box[i], cardBox.Box[j] = cardBox.Box[j], cardBox.Box[i]
	})
}

// 打印牌盒 (用于调试)
func (cardBox *CardBox) printCardBox() {
	for i := 0; i < len(cardBox.Box); i++ {
		fmt.Printf("%s ", cardBox.Box[i].String())
		if (i+1)%13 == 0 {
			fmt.Println()
		}
	}
}

// 获取一副新的扑克牌
func getNewCardBox() []Card {
	var cardBox []Card
	var colorSlice = getColorArray()
	var numberSlice = getNumberArray()
	for _, color := range colorSlice {
		for _, number := range numberSlice {
			cardBox = append(cardBox, Card{color, number})
		}
	}

	return cardBox
}
