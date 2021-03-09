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

// MinScore returns a score when Ace is equal to 1
func (h Hand) MinScore() int {
	s := 0
	for _, card := range h {
		s += min(int(card.Rank), 10)
	}
	return s
}

// Score returns a real score
func (h Hand) Score() int {
	s := h.MinScore()
	if s > 11 {
		return s
	}
	for _, card := range h {
		if card.Rank == deck.Ace {
			return s + 10
		}
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	// If dealer score <= 16, we hit
	// If dealer has a soft 17, we hit
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		var card deck.Card
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	playerScore, dealerScore := player.Score(), dealer.Score()
	fmt.Println("==FINAL SCORE==")
	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer)
	fmt.Println("Player score:", playerScore)
	fmt.Println("Dealer score:", dealerScore)
	switch {
	case playerScore > 21:
		fmt.Println("You busted")
	case dealerScore > 21:
		fmt.Println("Dealer busted")
	case playerScore > dealerScore:
		fmt.Println("You won!")
	case playerScore < dealerScore:
		fmt.Println("You lost")
	default:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
