package models

import (
	"log"
	"strings"
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

func (f Formation) Render() string {
    ROW_OFFSETS := []int{1, 1, 1, 1, 1, 1, 1}
    var rows []string
    for row := 0; row < len(f.formationSpots); row ++ {
        render := strings.Repeat(" ", ROW_OFFSETS[row])
        for col := 0; col < len(f.formationSpots[row]); col ++ {
            render += f.formationSpots[row][col].card.Render() + " "
        }
        rows = append(rows, render)
    }
    return strings.Join(rows, "\n")
}
