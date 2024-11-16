package models

import "log"

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

func (d *Formation) Init(cards []Card) Formation {
    formation := Formation{formationSpots: make([][]FormationSpot, 7)}
    for row := 0; row < 7; row ++ {
        formation.formationSpots[row] = make([]FormationSpot, row + 1)
        for col := 0; col <= row; col ++ {
            poppedCard := cards[0]
            cards = cards[1:]

            // TODO: setup coveredBy
            formationSpot := FormationSpot {card: &poppedCard, coveredBy: []Location{}}
            formation.formationSpots[row][col] = formationSpot
        }
    }

    if len(cards) != 0 {
        log.Fatalf("Expected 0 cards remaining")
    }
    return formation
}
