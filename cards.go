// package main

// import "fmt"

// type deck []string

// func main() {
// 	cards := newDeck()

// 	hand, remainingcards := deal(cards, 3)
// 	hand.print()
// 	remainingcards.print()
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
