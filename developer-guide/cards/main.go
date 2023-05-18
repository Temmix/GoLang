package main

func main() {
	cards := newDeckFromFile("my_cards") //newDeck()
	// hand, remainingCards := deal(cards, 5)
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()

}

func getCard() string {
	return "Testing"
}
