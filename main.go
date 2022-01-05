package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"

	//card := newCard()
	//fmt.Println(card)
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	//fmt.Println(cards)

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
