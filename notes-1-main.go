/*

Go is NOT OO.

*/

package main

import "fmt"

func main() {
	/* ------ Every declared variable must be used ------ */
	// declaring variables
	var myDeck string = "My deck"
	myOtherDeck := "My other deck"

	fmt.Println(myDeck, myOtherDeck)

	// declaring a slice
	cards:= []string{newCard(), newCard()}
	//adding items - creates a new slice and attributes it to the same variable
	cards = append(cards, "Six of diamonds")

	// iterating through a closed set.
	// range: used to iterate through every element of the set.
	// := is used bc for each iteration the previous index and item are disposed and redeclared 
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// a function must declare the type of it's return
func newCard() string {
	return "Ace of spades"
}

