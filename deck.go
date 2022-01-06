package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

/*
func (d deck) deal(number int) deck {
	cards := deck{}
	num := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		num = rand.Intn(52-0) + 0
		cards = append(cards, d[num])
	}

	return cards
}
*/

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
