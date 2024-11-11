package main

import(
    "fmt"
    "os"
    "pyramid-solitaire/models"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    // formation Formation
    drawPile []models.Card
    discardPile []models.Card
    // livePile []models.Card (nullable)
}

func initialModel() model {
    return model{
        // formation Formation
        // drawPile: []models.Card{},
        discardPile: []models.Card{},
        // livePile []models.Card (nullable)
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
    s := "Press q to exit"
    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error: %v", err)
        os.Exit(1)
    }
}
