package main

func main() {
	// cards := newDeck()
	// cards := newDeckFromFile("daenam_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
