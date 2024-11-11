package models

type Deck struct {
    cards []Card
}

func (d Deck) Init() Deck {
    var deck Deck

    for suit := Clubs; suit <= Spades; suit++ {
        for rank := Ace; rank <= King; rank++ {
            deck.cards = append(deck.cards, Card{Rank: rank, Suit: suit})
        }
    }

    return deck
}

// TODO: shuffle

// TODO: draw
