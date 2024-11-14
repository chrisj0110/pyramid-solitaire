package models

type Location struct {
    row int
    col int
}

type FormationSpot struct {
    card *Card // nil means there is no card here
    coveredBy []Location
}
