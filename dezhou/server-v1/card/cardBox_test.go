package card

import "testing"

func TestNewCardBox(t *testing.T) {
	cardBox := NewCardBox()
	cardBox.printCardBox()
}

func TestCardBox_ShuffleCardBox(t *testing.T) {
	cardBox := NewCardBox()
	cardBox.ShuffleCardBox()
	cardBox.printCardBox()
}

func TestCardBox_GetOneCard(t *testing.T) {
	cardBox := NewCardBox()
	for i := 0; i < 10; i++ {
		t.Log(cardBox.GetOneCard())
	}
}
