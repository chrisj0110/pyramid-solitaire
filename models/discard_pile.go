package models

import (
    "fmt"
    "strings"
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

func (dp *DiscardPile) SelectCard() {
    if len(dp.cards) == 0 {
        return
    }
    dp.cards[len(dp.cards)-1].selected = true
}

func (dp *DiscardPile) UnselectCard() {
    // just unselect all of them
    for idx := 0; idx < len(dp.cards); idx++ {
        dp.cards[idx].selected = false
    }
}

func (dp *DiscardPile) RemoveSelectedCards() {
    // build a new list, don't include cards to be removed
    newCardList := []Card{}

    // just go through all of them
    for idx := 0; idx < len(dp.cards); idx++ {
        if dp.cards[idx].selected {
            dp.cards[idx].selected = false
        } else {
            newCardList = append(newCardList, dp.cards[idx])
        }
    }
    dp.cards = newCardList
}

func (dp DiscardPile) GetSelectedCards() []Card {
    // just go through all of them
    cards := []Card{}
    for idx := 0; idx < len(dp.cards); idx++ {
        if dp.cards[idx].selected {
            cards = append(cards, dp.cards[idx])
        }
    }
    return cards
}

func (dp DiscardPile) Render() string {
    cardLen := len(dp.cards)

    prefix := ""
    cardsToShow := []Card{}
    if cardLen == 0 {
        prefix = "No cards"
    } else if cardLen > 5 {
        // if 7 cards, show indexes 2:6
        // if 8 cards, show indexes 3:7
        cardsToShow = dp.cards[cardLen-5:cardLen]
        prefix = fmt.Sprintf("Last 5 of %v cards: ", cardLen)
    } else {
        prefix = "All cards: "
        cardsToShow = dp.cards
    }

    cardStrs := []string{}
    for i := 0; i < len(cardsToShow); i++ {
        cardStrs = append(cardStrs, cardsToShow[i].Render())
    }

    return fmt.Sprintf("%v%v", prefix, strings.Join(cardStrs, " "))
}
