package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
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

/*
func saveToFile(d deck, file string) {
	err := ioutil.WriteFile(file, []byte(d.toString()), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
*/

func (d deck) saveToFile(file string) error {
	return ioutil.WriteFile(file, []byte(d.toString()), 0666)
}

func readFromFile(file string) deck {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		// also can use os.Exit(1) to stop the program
		log.Println(err, "\n Craeting new deck")
		content = []byte(newDeck().toString())
	}
	return deck(strings.Split(string(content), ","))
}

func (d deck) shuffle() {
	//rand.Seed(time.Now().UnixNano())

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for index := range d {
		num := r.Intn(51-0) + 0
		// d[index], d[number] = d[number], d[index] fancy swapping
		placeholder := d[index]
		d[index] = d[num]
		d[num] = placeholder
	}
}
