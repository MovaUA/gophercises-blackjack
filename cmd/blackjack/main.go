package main

import (
	"fmt"
	"strings"

	"github.com/movaua/gophercises-deck/pkg/deck"
)

// Hand is a slice of cards in the hand
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := 0; i < len(h); i++ {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString prints cards of a dealer
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var player Hand
	var dealer Hand

	for i := 0; i < 2; i++ {
		for _, h := range []*Hand{&player, &dealer} {
			var card deck.Card
			card, cards = draw(cards)
			*h = append(*h, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)stand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			var card deck.Card
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	fmt.Println("==FINAL SCORE==")
	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
