package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	var cards map[rune][]string = make(map[rune][]string)
	for _, suit := range suits {
		var ranksOfSuit []string = make([]string, 9)
		for i, rank := range ranks {
			ranksOfSuit[i] = fmt.Sprintf("%c%c", rank, suit)
		}
		cards[suit] = ranksOfSuit
	}

	for _, suit := range suits {
		for _, rank := range cards[suit] {
			fmt.Print(rank + "\t")
		}
		fmt.Println()
	}
}
