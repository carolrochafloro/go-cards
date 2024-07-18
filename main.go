package main

func main() {
	cards := newDeck()

	// functions declared in the same package don't have to be imported
	// assingning two different returns of a function to two different variables
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
