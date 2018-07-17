package deck

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cards := New()
	expected := []Card{
		Card{
			Rank: ace,
			Suit: clubs,
		},
		Card{
			Rank: two,
			Suit: clubs,
		},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{
			Rank: ace,
			Suit: diamonds,
		},
	}
	if cards[0] != expected[0] || cards[1] != expected[1] || cards[13] != expected[13] {
		t.Fatalf("Got:\n%v\nWant:\n%v\n", cards, expected)
	}

	deck1 := New(Shuffle())
	deck2 := New(Shuffle())
	if reflect.DeepEqual(deck1, deck2) {
		t.Fatalf("Got two the same decks")
	}

	cards = New(MultiDeck(1))
	if len(cards) != 52 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(MultiDeck(3))
	if len(cards) != 156 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(Remove(ace, two, three, four))
	if len(cards) != 36 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(Remove(king, jack), MultiDeck(2))
	if len(cards) != 88 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(SortByRank(func(i, j rank) bool {
		return i < j
	}))
	if cards[0].Rank != ace {
		t.Fatalf("Got=\n%v\nExpected that first rank will be %v", cards, ace.String())
	}

	cards = New(SortByRank(func(i, j rank) bool {
		return i > j
	}))
	if cards[0].Rank != king {
		t.Fatalf("Got=\n%v\nExpected that first rank will be %v", cards, king.String())
	}

	cards = New(SortBySuit(func(i, j suit) bool {
		return i > j
	}))
	if cards[0].Suit != spades {
		t.Fatalf("Got=\n%v\nExpected that first suit will be %v", cards, spades.String())
	}

	cards = New(SortByRank(func(i, j rank) bool {
		return i > j
	}), SortBySuit(func(i, j suit) bool {
		return i > j
	}))
	if cards[0].Rank != king && cards[0].Suit != spades {
		t.Fatalf("Got=\n%v\nExpected that first card will be %v %v", cards, king.String(), spades.String())
	}
}

func TestRankString(t *testing.T) {
	a := ace.String()
	expected := "ace"
	if a != expected {
		t.Fatalf("Got: %v, Expected: %v", a, expected)
	}

	var f rank = -2
	expected = "rank(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}

func TestSuitString(t *testing.T) {
	s := spades.String()
	expected := "spades"
	if s != expected {
		t.Fatalf("Got: %v, Expected: %v", s, expected)
	}

	var f suit = -2
	expected = "suit(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}
