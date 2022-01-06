package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"

	//card := newCard()
	//fmt.Println(card)
	cards := newDeck()
	//hand, remainingCard := deal(cards, 15)
	//fmt.Println(cards)
	//fmt.Println(cards.toString())
	//fmt.Println(cards)
	cards.saveToFile("deck.tx")
	c := readFromFile("deck.txtt")
	c.shuffle()
	c.print()

	//hand.print()
	//fmt.Println("Now the remaining cards")
	//remainingCard.print()
	// Type coversion
	/*
	   greeting := "Hi there!"
	   	fmt.Println([]byte(greeting))
	*/
}
