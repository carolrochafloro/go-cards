package main

import "fmt"

/*

Creating a new type deck using an existing type and adding more funcionalities,
such as functions with a receiver.

*/

type deck []string 


func newDeck() deck {

	cards := deck{}
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	// the indexes where replaced with underscore bc they won't be used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards

}

/*
- deck is the receiver and any variable of this type gets access to this function.
- by convention the receiver is usually named with a one or two
letter abbreviation that matches it's type
- it's equivalent to "this" statement in other langs
*/ 
func(d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returning 2 values from a function and receiving parameters
func deal(d deck, handSize int) (deck, deck) {
	/*
	
	slice range syntax:
	sliceName[startIndexIncluded:endIndexNotIncluded]

	abbreviation:
	sliceName[:endIndex] => from the beginning to the end index (not included)
	slineName[startIndex:] => from the start index to the end (all included)

	*/
	return d[:handSize], d[handSize:]
}