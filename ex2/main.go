package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	if zodiacSign == Aries {
		fmt.Println("21.03. - 20.04")
	} else if zodiacSign == Taurus {
		fmt.Println("21.04. - 21.05")
	} else {
		fmt.Println("")
	}
	// TODO: ...and consider a default case.
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
