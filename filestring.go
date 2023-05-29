// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strings"
// )

// type deck []string

// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("That's a file")
// }

// func newDeck() deck {
// 	cards := deck{}

// 	cardSuits := []string{"Spade", "Diamond", "Club", "Heart"}
// 	cardValue := []string{"One", "two", "Ace", "three"}

// 	for _, suit := range cardSuits {
// 		for _, val := range cardValue {
// 			cards = append(cards, suit+" of "+val)
// 		}
// 	}
// 	return cards
// }

// func (d deck) print() {
// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }

// func deal(d deck, handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

// func (d deck) saveToFile(filename string) error {
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }
