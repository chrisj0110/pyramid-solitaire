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
    return tea.ClearScreen
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

    // TODO: this is just for testing
    // view += fmt.Sprintf("\n%v cards remaining in deck", m.deck.GetRemainingCount())
    // if len(m.discardPile) > 0 {
    //     view += fmt.Sprintf("\n%v top card in discard pile", m.discardPile[len(m.discardPile)-1].Render())
    // }

    return view
}

// TODO: why doesn't this work?
// go get github.com/charmbracelet/bubbles/table
// func legendRender() string {
//     columns := []table.Column{
//         {Title:"Key"},
//         {Title:"Action"},
//     }
//     rows := []table.Row{
//         {"asdfjkl", "select from formation"},
//         {"r", "refresh"},
//         {"n", "next card"},
//         // p - play from discard pile
//         // esc - unselect card
//         // u - undo
//         {"q", "quit"},
//     }
//     t := table.New(table.WithColumns(columns), table.WithRows(rows))
//     return t.View()
// }

func legendRender() string {
    return "asdfjkl - select from formation\n" +
    "r - refresh\n" +
    "n - next card\n" +
    // p - play from discard pile
    // esc - unselect card
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
