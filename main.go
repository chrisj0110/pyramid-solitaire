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
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "n":
            card, err := m.deck.Draw()
            if err != nil {
                // TODO: set an error message in the model to display
                log.Fatalf("no cards left")
            }
            m.discardPile.Add(card)

            return m, nil
        }
    }

    return m, nil
}

func (m model) View() string {
    contentSquareStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.RoundedBorder())
    titleStyle := lipgloss.NewStyle().Bold(true)

    // formation
    view := lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Pyramid "), contentSquareStyle.Render(m.formation.Render()))
    view += "\n"

    // discard pile
    view += lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(" Discard Pile "), contentSquareStyle.Render(m.discardPile.Render()))
    view += "\n"

    // TODO: this is just for testing
    // view += fmt.Sprintf("\n%v cards remaining in deck", m.deck.GetRemainingCount())
    // if len(m.discardPile) > 0 {
    //     view += fmt.Sprintf("\n%v top card in discard pile", m.discardPile[len(m.discardPile)-1].Render())
    // }

    return view
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error: %v", err)
        os.Exit(1)
    }
}
