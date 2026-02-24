package main

import (
	"bytes"
	"fmt"
	"os"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	questions     []string
	inputField    textinput.Model
	currentActive int
	answers       []string
}

func InitialModel() model {

	questions := []string{
		"what is your name",
		"what is your age",
		"what do u do ",
	}

	t := textinput.New()
	t.Placeholder = "type something"
	t.Focus()

	return model{
		questions:     questions,
		inputField:    t,
		answers:       []string{},
		currentActive: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			return m, tea.Quit
		case "enter":
			m.answers = append(m.answers, m.inputField.Value())
			m.inputField.SetValue("")
			m.currentActive++

			// check last question
			if m.currentActive >= len(m.questions) {
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.inputField, cmd = m.inputField.Update(msg)

	return m, cmd
}

func (m model) View() string {

	var temp bytes.Buffer

	if m.currentActive >= len(m.questions) {
		// show all answers
		for i := range m.questions {
			line := fmt.Sprintf(
				"%s: %s\n",
				m.questions[i],
				m.answers[i],
			)
			temp.WriteString(line)
		}

		temp.WriteString("\nPress q to quit.")
		return temp.String()
	}

	q := m.questions[m.currentActive]
	return fmt.Sprintf("%s\n\n%s\n ", q, m.inputField.View())

}

func main() {

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
