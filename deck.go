package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*

Creating a new type deck using an existing type and adding more funcionalities,
such as functions with a receiver.

*/

type deck []string 


func newDeck() deck {

	cards := deck{}
	cardValues := []string {"Ace", "Two", "Three", "Four"}
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}


	// the indexes where replaced with underscore bc they won't be used
	for _, suit := range cardValues {

		for _, value := range cardSuits {

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

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
	
}

// using "os" package to write a file, converting the string deck to byte
func (d deck) saveToFile(fileName string) error {

	return os.WriteFile(fileName, []byte(d.toString()), 0666)

}

func deckFromFile(fileName string) deck {

	// bs = byte slice
	bs, err := os.ReadFile(fileName)

	// error handling - log and break
	if err != nil {

		fmt.Println(" ==== Error: ", err)
		os.Exit(1)

	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {

	// generating a seed to the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {

		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}

