package main

//"net/http"

func main() {
	//cards := newDeck()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	//cards.saveToFile("my_cards")

	//cards := newDeck()
	//cards.shuffle()
	//cards.print()
	//fmt.Println("----------")
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	//cards.print()
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.print()
	//hand, remainingDeck := deal(cards, 5)
	//fmt.Println("Hand:")
	//hand.print()
	//fmt.Println("Remaining Deck:")
	//remainingDeck.print()
}
