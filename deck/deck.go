package deck

import (
	"math/rand"
	"time"
)

type suit int

const (
	clubs suit = iota
	diamonds
	hearts
	spades
)

type rank int

const (
	ace rank = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
)

// Card is representing abstract playing card
type Card struct {
	Rank rank
	Suit suit
}

const (
	cardsForOneDeck = 52
	cardsForOneSuit = 13
	allSuits        = 4
)

// Deck is TODO
type Deck func([]Card) []Card

// New returns deck TODO
func New(opts ...Deck) []Card {
	cards := make([]Card, 0, cardsForOneDeck)
	for i := 0; i < allSuits; i++ {
		for j := 1; j <= cardsForOneSuit; j++ {
			cards = append(cards, Card{
				Rank: rank(j),
				Suit: suit(i),
			})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// Shuffle the deck
func Shuffle() Deck {
	return func(cards []Card) []Card {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

// MultiDeck appends q decks to original
func MultiDeck(q int) Deck {
	if q <= 1 {
		return func(cards []Card) []Card { return cards }
	}
	return func(cards []Card) []Card {
		tmp := make([]Card, len(cards), q*len(cards))
		copy(tmp, cards)
		for q != 1 {
			tmp = append(tmp, cards...)
			q--
		}
		return tmp
	}
}
