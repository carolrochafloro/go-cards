package main

func main() {
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Queen of spades")

	cards.print()

}

func newCard() string {
	return "Five of diamonds"
}