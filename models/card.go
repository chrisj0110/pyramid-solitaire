package models

import "github.com/charmbracelet/lipgloss"

type Card struct {
    Rank CardRank
    Suit CardSuit
    selected bool
}

var black = lipgloss.Color("0")
var red = lipgloss.Color("9")
var yellow = lipgloss.Color("11")
var white = lipgloss.Color("15")
var redCard = lipgloss.NewStyle().Foreground(red).Background(white)
var redSelectedCard = lipgloss.NewStyle().Foreground(red).Background(yellow)
var blackCard = lipgloss.NewStyle().Foreground(black).Background(white)
var blackSelectedCard = lipgloss.NewStyle().Foreground(black).Background(yellow)

func (c Card) Render() string {
    content := " " + c.Rank.String() + c.Suit.String() + " "
    if c.Suit.isRed() {
        if c.selected {
            return redSelectedCard.Render(content)
        } else {
            return redCard.Render(content)
        }
    } else {
        if c.selected {
            return blackSelectedCard.Render(content)
        } else {
            return blackCard.Render(content)
        }
    }
}

type CardRank int

const (
    Ace CardRank = iota + 1
    Two
    Three
    Four
    Five
    Six
    Seven
    Eight
    Nine
    Ten
    Jack
    Queen
    King
)

func (r CardRank) String() string {
    return [...]string{" A", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9", "10", " J", " Q", " K"}[r-1]
}

type CardSuit int

const (
    Clubs CardSuit = iota
    Diamonds
    Hearts
    Spades
)

func (s CardSuit) String() string {
    return [...]string{"♣️", "♦️", "❤️", "♠️"}[s]
}

func (s CardSuit) isRed() bool {
    if s == 0 || s == 3 {
        return false
    }
    return true
}
