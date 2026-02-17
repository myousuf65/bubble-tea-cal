package counter

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func initialModel() model {
	return model{
		count: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "up":
			m.count++
		case "down":
			m.count--
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		return m, nil

	case tea.WindowSizeMsg:
		fmt.Println("window sizing message here", msg.Width, msg.Height)
	}

	return m, nil
}

func (m model) View() string {
	output := fmt.Sprintf("count: %d", m.count)

	return output
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
