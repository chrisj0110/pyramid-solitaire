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
    discardPile []models.Card
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
    return model{
        formation: formation.Init(formationCards),
        deck: deck,
        discardPile: []models.Card{},
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
        }
    }

    return m, nil
}

func (m model) View() string {
    formationStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.RoundedBorder())
    formationStr := m.formation.Render()

    view := formationStyle.Render(formationStr)
    return view
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error: %v", err)
        os.Exit(1)
    }
}
