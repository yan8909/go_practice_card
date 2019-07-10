package main

func main() {

	//var card string = "Ace of Spades"
	cards := newDeck()

	hand, reaminingCards := deal(cards, 5)

	hand.print()
	reaminingCards.print()
}
