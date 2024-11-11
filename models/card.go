package models

type Card struct {
    Rank CardRank
    Suit CardSuit
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
    return [...]string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}[r-1]
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
