package models

import (
	"log"
	"strings"
    "github.com/charmbracelet/lipgloss"
)

type Location struct {
    row int
    col int
}

type FormationSpot struct {
    card *Card // nil means there is no card here
    coveredBy []Location
}

type Formation struct {
    formationSpots [][]FormationSpot
}

func (f *Formation) Init(cards []Card) Formation {
    formation := Formation{formationSpots: make([][]FormationSpot, 7)}
    for row := 6; row >= 0; row -- {
        formation.formationSpots[row] = make([]FormationSpot, row + 1)
        for col := 0; col <= row; col ++ {
            poppedCard := cards[0]
            cards = cards[1:]

            coveredBy := []Location{}
            if row < 6 {
                coveredBy = []Location{
                    { row: row+1, col: col },
                    { row: row+1, col: col+1 },
                }
            }
            formation.formationSpots[row][col] = FormationSpot {
                card: &poppedCard,
                coveredBy: coveredBy,
            }
        }
    }

    if len(cards) != 0 {
        log.Fatalf("Expected 0 cards remaining")
    }
    return formation
}

func (f *Formation) SelectCard(idx int) {
    // find the lowest card in the formation and select it
    for row := 6; row >= 0; row-- {
        if f.formationSpots[row][idx].card != nil {
            f.formationSpots[row][idx].card.selected = true
            return
        }
    }
}

func (f *Formation) UnselectCard() {
    // just unselect all of them
    for row := 0; row < len(f.formationSpots); row++ {
        for col := 0; col < len(f.formationSpots[row]); col++ {
            if f.formationSpots[row][col].card != nil {
                f.formationSpots[row][col].card.selected = false
            }
        }
    }
}

func (f *Formation) RemoveSelectedCards() {
    // just go through all of them
    for row := 0; row < len(f.formationSpots); row++ {
        for col := 0; col < len(f.formationSpots[row]); col++ {
            if f.formationSpots[row][col].card != nil && f.formationSpots[row][col].card.selected {
                f.formationSpots[row][col].card.selected = false
                f.formationSpots[row][col].card = nil
            }
        }
    }
}

func (f Formation) GetSelectedCards() []Card {
    // just go through all of them
    cards := []Card{}
    for row := 0; row < len(f.formationSpots); row++ {
        for col := 0; col < len(f.formationSpots[row]); col++ {
            if f.formationSpots[row][col].card != nil && f.formationSpots[row][col].card.selected {
                cards = append(cards, *f.formationSpots[row][col].card)
            }
        }
    }
    return cards
}

func (f Formation) Render() string {
    trueBlack := lipgloss.Color("#000000")
    emptySpace := lipgloss.NewStyle().Background(trueBlack)

    ROW_OFFSETS := []int{1, 1, 1, 1, 1, 1, 1}
    var rows []string
    for row := 0; row < len(f.formationSpots); row ++ {
        render := strings.Repeat(emptySpace.Render(" "), ROW_OFFSETS[row])
        for col := 0; col < len(f.formationSpots[row]); col ++ {
            if f.formationSpots[row][col].card == nil {
                // if card is nil, then render an empty card
                render += RenderEmptySpot() + emptySpace.Render(" ")
            } else {
                render += f.formationSpots[row][col].card.Render() + emptySpace.Render(" ")
            }
        }
        rows = append(rows, render)
    }
    return strings.Join(rows, "\n")
}
