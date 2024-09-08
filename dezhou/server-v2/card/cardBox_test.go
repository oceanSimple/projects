package card

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCardBox(t *testing.T) {
	box := NewCardBox()
	box.printCardBox()
}

func TestCardBox_ShuffleCardBox(t *testing.T) {
	box := NewCardBox()
	box.ShuffleCardBox()
	box.printCardBox()
}

func TestCardBox_GetOneCard(t *testing.T) {
	box := NewCardBox()
	box.ShuffleCardBox()

	c1 := box.GetOneCard()
	c2 := box.GetOneCard()
	c3 := box.GetOneCard()
	fmt.Printf("%s %s %s\n", c1.String(), c2.String(), c3.String())
	assert.Equal(t, 49, len(box.Box))
}
