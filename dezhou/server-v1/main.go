package main

import (
	"fmt"
	"server-v1/card"
	"server-v1/output"
	"server-v1/pattern"
	"server-v1/tool"
)

var cardMap = make(map[string]int)

func main() {
	// Start the test server
	// print some information
	Tips()

	for {
		fmt.Println(output.Info() + "Start a new round")
		// var uuid = uuid.New().String()
		var cards = GetCards()
		weight, _ := pattern.Judge(cards)

		// print the cards
		fmt.Println(output.Default() + "The cards are:")
		for _, c := range cards {
			fmt.Printf("%s ", c.DyeString())
		}
		fmt.Println()

		// wait for the user to enter the weight
		var userInput int

		for {
			fmt.Print(output.Default() + "Please enter the weight: ")
			_, err := fmt.Scanf("%d", &userInput)
			fmt.Println()
			if err != nil {
				// fmt.Println("Invalid input, please enter again")
				continue
			}
			break
		}

		// compare the weight and user input
		if userInput == weight {
			fmt.Println(output.Default() + "Congratulations! You are right!")
		} else {
			fmt.Println(output.Error() + "Sorry, you or system may have a mistake")
			fmt.Println(output.Error() + "Please put again: ")
			for {
				fmt.Print(output.Default() + "Please enter the weight: ")
				_, err := fmt.Scanf("%d", &userInput)
				fmt.Println()
				if err != nil {
					continue
				}
				break
			}
			if userInput == weight {
				fmt.Println(output.Default() + "Congratulations! You are right!")
			} else {
				fmt.Println(output.Error() + "The answers are still not the same, we will note this err! Thanks!")
			}
		}
		fmt.Println("=====================================")
		fmt.Println("=====================================")
	}
}

// Tips print
func Tips() {
	fmt.Println(output.Info() + "Welcome to the test server!")
	fmt.Println(output.Info() + "The server will give you 7 cards")
	fmt.Println(output.Info() + "The only thing you need to do is judge weight")
	fmt.Println("=====================================")
	fmt.Println(output.DyeText("8-同花顺，7-四条，6-葫芦，5-同花，4-顺子，3-三条，2-两对，1-一对，0-散牌", output.Green))
	fmt.Println("=====================================")
}

func GetCards() []*card.Card {
	var cards []*card.Card
	box := card.NewCardBox()
	box.ShuffleCardBox()
	for i := 0; i < 7; i++ {
		cards = append(cards, box.GetOneCard())
	}
	tool.SortCardsByNumber(cards)
	return cards
}
