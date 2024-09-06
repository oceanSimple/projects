package main

import (
	"fmt"
	"testing"
)

func TestGetCards(t *testing.T) {
	cards := GetCards()
	for i := 0; i < len(cards); i++ {
		fmt.Printf("%s ", cards[i].DyeString())
	}
}
