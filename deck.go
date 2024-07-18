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

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	

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