package main

func main() {
	cards := newDeck()
	cards.randomGenerator()
	cards.printCards()
}
