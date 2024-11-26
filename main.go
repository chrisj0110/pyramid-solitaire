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
    }
}

func (m model) Init() tea.Cmd {
    return tea.ClearScreen
}

// 0-6 for diagonals in the formation; 7 for discard pile
func (m* model) selectCard(idx int) {
    if idx < 7 {
        m.formation.SelectCard(idx)
    } else {
        m.discardPile.SelectCard()
    }
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
        case "ctrl+c", "q":
            return m, tea.Quit
        case "a":
            m.selectCard(0)
            m.tryPlayCards()
            return m, nil
        case "s":
            m.selectCard(1)
            m.tryPlayCards()
            return m, nil
        case "d":
            m.selectCard(2)
            m.tryPlayCards()
            return m, nil
        case "f":
            m.selectCard(3)
            m.tryPlayCards()
            return m, nil
        case "j":
            m.selectCard(4)
            m.tryPlayCards()
            return m, nil
        case "k":
            m.selectCard(5)
            m.tryPlayCards()
            return m, nil
        case "l":
            m.selectCard(6)
            m.tryPlayCards()
            return m, nil
        case "p":
            m.selectCard(7)
            m.tryPlayCards()
            return m, nil
        case "c":
            m.unselectCard()
            return m, nil
        case "n":
            card, err := m.deck.Draw()
            if err != nil {
                // TODO: set an error message in the model to display
                log.Fatalf("no cards left")
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

    contentSquareStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.RoundedBorder()).Width(WIDTH)
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
    return "asdfjkl - select from formation\n" +
    "r - refresh\n" +
    "n - next card\n" +
    "p - play from discard pile\n" +
    "c - change mind\n" +
    // u - undo
    "q - quit"
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error: %v", err)
        os.Exit(1)
    }
}
