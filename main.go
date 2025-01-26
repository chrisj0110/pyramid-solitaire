package main

import (
	"fmt"
	"log"
	"os"
	"pyramid-solitaire/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
    formation models.Formation
    deck models.Deck
    discardPile models.DiscardPile
    message string
    selectedRank *models.CardRank
}

func initialModel() model {
    var deck models.Deck
    deck = deck.Init()
    deck.Shuffle()

    var formationCards []models.Card
    for i := 1; i <= 28; i++ {
        card, err := deck.Draw()
        if err != nil {
            log.Fatalf("Error: %v", err)
        }
        formationCards = append(formationCards, card)
    }
    if deck.GetRemainingCount() != 52 - 28 {
        log.Fatalf("Expected 24 cards remaining")
    }

    var formation models.Formation
    var discardPile models.DiscardPile
    return model{
        formation: formation.Init(formationCards),
        deck: deck,
        discardPile: discardPile.Init(),
        message: "",
        selectedRank: nil,
    }
}

func (m model) Init() tea.Cmd {
    return tea.ClearScreen
}

func (m* model) selectCard(cardRank models.CardRank, cardSuit models.CardSuit) error {
    if !m.formation.SelectCard(cardRank, cardSuit) && !m.discardPile.SelectCard(cardRank, cardSuit) {
        return fmt.Errorf("no card found to play")
    }
    return nil
}

func (m* model) unselectCard() {
    m.formation.UnselectCard()
    m.discardPile.UnselectCard()
}

func (m* model) removeSelectedCards() {
    m.formation.RemoveSelectedCards()
    m.discardPile.RemoveSelectedCards()
}

func (m model) getSelectedCards() []models.Card {
    cards := m.formation.GetSelectedCards()
    return append(cards, m.discardPile.GetSelectedCards()...)
}

func (m* model) tryPlayCards() {
    cards := m.getSelectedCards()
    if len(cards) == 1 && int(cards[0].Rank) == 13 {
        m.removeSelectedCards()
    } else if len(cards) == 2 && int(cards[0].Rank) + int(cards[1].Rank) == 13 {
        m.removeSelectedCards()
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.message = ""
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "Q":
            return m, tea.Quit
        case "G":
            m = initialModel()
            return m, nil
        case "a", "2", "3", "4", "5", "6", "7", "8", "9", "t", "j", "q", "k":
            m.selectedRank = models.CardRankFromString(msg.String())
            if m.selectedRank == nil {
                log.Fatalf("card rank not found")
            }

            m.message = "now select suit"
            return m, nil
        case "c", "d", "h", "s":
            if m.selectedRank == nil {
                m.message = "you need to choose a card rank before a card suit"
                return m, nil
            }

            selectedSuit := models.CardSuitFromString(msg.String())
            if selectedSuit == nil {
                log.Fatalf("card suit not found")
            }
            err := m.selectCard(*m.selectedRank, *selectedSuit)
            if err != nil {
                m.message = fmt.Sprintf("%v", err)
                return m, nil
            }

            m.tryPlayCards()
            if m.formation.IsGameOver() {
                m.message = "YOU WIN!!!"
                return m, tea.Quit
            }
            return m, nil
        case "u":
            // TODO: or if card isn't selected, undo the last turn
            m.unselectCard()
            return m, nil
        case "n":
            card, err := m.deck.Draw()
            if err != nil {
                m.message = "No cards left!"
                return m, nil
            }
            m.discardPile.Add(card)

            m.unselectCard()

            return m, nil
        case "r":
            return m, tea.ClearScreen
        }
    }

    return m, nil
}

func (m model) View() string {
    WIDTH := 56 // keep this as an even number
    // example of how to calculate width:
    // width := lipgloss.Width(contentSquareStyle.Render(m.discardPile.Render()))

    trueBlack := lipgloss.Color("#000000")
    contentSquareStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.RoundedBorder()).Width(WIDTH).Background(trueBlack)
    titleStyle := lipgloss.NewStyle().Bold(true)

    // formation
    view := lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Pyramid "), contentSquareStyle.Render(m.formation.Render()))
    view += "\n"

    // discard pile
    view += lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Discard Pile "), contentSquareStyle.Render(m.discardPile.Render()))
    view += "\n"

    // deck
    view += lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Deck "), contentSquareStyle.Render(m.deck.Render()))
    view += "\n"

    // legend
    view += lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Legend "), contentSquareStyle.Render(legendRender()))
    view += "\n"

    // message
    view += lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Message "), contentSquareStyle.Render(m.message))
    view += "\n"

    return view
}

func legendRender() string {
    return "a, 2-9, t, j, q, k - select from formation\n" +
    "c - clubs, d - diamonds, h - hearts, s - spades\n" +
    "r - refresh UI\n" +
    "n - next card\n" +
    "p - play from discard pile\n" +
    "u - undo\n" +
    "G - new game\n" +
    "Q - quit"
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error: %v", err)
        os.Exit(1)
    }
}
