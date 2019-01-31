package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("daenam_cards")
	cards.print()
}
