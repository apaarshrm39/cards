package main

import (
	"fmt"
)

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"

	//card := newCard()
	//fmt.Println(card)
	cards := newDeck()
	hand, remainingCard := deal(cards, 15)
	//fmt.Println(cards)
	fmt.Println(cards.toString())

	hand.print()
	fmt.Println("Now the remaining cards")
	remainingCard.print()
	// Type coversion
	/*
	   greeting := "Hi there!"
	   	fmt.Println([]byte(greeting))
	*/
}
