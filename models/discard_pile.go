package models

import (
    "fmt"
)

type DiscardPile struct {
    cards []Card
}

func (dp *DiscardPile) Init() DiscardPile {
    return DiscardPile{cards: []Card{}}
}

func (dp *DiscardPile) Add(card Card) {
    dp.cards = append(dp.cards, card)
}

func (dp DiscardPile) Render() string {
    return fmt.Sprintf("%v cards", len(dp.cards))
}
