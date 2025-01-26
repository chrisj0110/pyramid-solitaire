package models

import (
    "github.com/charmbracelet/lipgloss"
)

type Card struct {
    Rank CardRank
    Suit CardSuit
    selected bool
}

var trueBlack = lipgloss.Color("#000000")
var black = lipgloss.Color("#1e1e2e")
var red = lipgloss.Color("#f38ba8")
var yellow = lipgloss.Color("#f9e2af")
var white = lipgloss.Color("#f5e0dc")
var redCard = lipgloss.NewStyle().Foreground(red).Background(white).Bold(true)
var redSelectedCard = lipgloss.NewStyle().Foreground(red).Background(yellow).Bold(true)
var blackCard = lipgloss.NewStyle().Foreground(black).Background(white).Bold(true)
var blackSelectedCard = lipgloss.NewStyle().Foreground(black).Background(yellow).Bold(true)
var emptySpot = lipgloss.NewStyle().Background(trueBlack)

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

func RenderEmptySpot() string {
    content := " " + " " + "   " + " "
    return emptySpot.Render(content)
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

var commandToCardRank = map[string]CardRank{
    "a": Ace,
    "2": Two,
    "3": Three,
    "4": Four,
    "5": Five,
    "6": Six,
    "7": Seven,
    "8": Eight,
    "9": Nine,
    "t": Ten,
    "j": Jack,
    "q": Queen,
    "k": King,
}

func CardRankFromString(menuCommand string) *CardRank {
    if rank, ok := commandToCardRank[menuCommand]; ok {
        return &rank
    }
    return nil
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

var commandToCardSuit = map[string]CardSuit{
    "c": Clubs,
    "d": Diamonds,
    "h": Hearts,
    "s": Spades,
}

func CardSuitFromString(menuCommand string) *CardSuit {
    if suit, ok := commandToCardSuit[menuCommand]; ok {
        return &suit
    }
    return nil
}

