package models

import (
    "fmt"
    "math/rand"
    "time"
    "log"
    "github.com/charmbracelet/lipgloss"
)

type Deck struct {
    cards []Card
}

var darkGreen = lipgloss.Color("22")
var cardBack = lipgloss.NewStyle().Foreground(white).Background(darkGreen)

func (d *Deck) Init() Deck {
    var deck Deck

    for suit := Clubs; suit <= Spades; suit++ {
        for rank := Ace; rank <= King; rank++ {
            deck.cards = append(deck.cards, Card{Rank: rank, Suit: suit})
        }
    }

    if deck.GetRemainingCount() != 52 {
        log.Fatalf("Deck should be 52")
    }

    return deck
}

func (d Deck) Render() string {
    numberRemaining := d.GetRemainingCount()
    numberRemainingStr := fmt.Sprintf("%v", numberRemaining)
    if numberRemaining < 10 {
        numberRemainingStr = fmt.Sprintf(" %v", numberRemaining)
    }
    content := " " + numberRemainingStr + " "

    return cardBack.Render(content)
}

func (d *Deck) Shuffle() {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    rng.Shuffle(len(d.cards), func(i, j int) {
        d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
    })
}

func (d *Deck) Draw() (Card, error) {
    if len(d.cards) == 0 {
        return Card{}, fmt.Errorf("deck is empty")
    }

    card := d.cards[0]
    d.cards = d.cards[1:]

    return card, nil
}

func (d Deck) GetRemainingCount() int {
    return len(d.cards)
}
