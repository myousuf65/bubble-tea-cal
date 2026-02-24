package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dates       map[string][]string
	activeDate  int
	activeMonth string
}

func InitialModel() model {

	m := make(map[string][]string)

	m["January"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["February"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28"}

	m["March"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["April"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}

	m["May"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["June"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}

	m["July"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["August"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["September"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}

	m["October"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	m["November"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}

	m["December"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	// get todays date
	day := time.Now().Day()
	month := time.Now().Month()

	return model{
		dates:       m,
		activeDate:  day,
		activeMonth: month.String(),
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

		case "left":
			m.activeDate--
			return m, nil

		case "right":
			m.activeDate++
			return m, nil

		case "up":
			m.activeDate -= 7
			return m, nil

		case "down":
			m.activeDate += 7
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {

	// user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// filename
	filename := strconv.Itoa(m.activeDate) + ".txt"

	// directory
	filePath := filepath.Join(home, "journal", m.activeMonth, filename)


	

	var output strings.Builder
	output.WriteString("currentActive: " + strconv.Itoa(m.activeDate) + "\n")
	counter := 0
	// ------------current month ----------------
	for _, v := range m.dates[m.activeMonth] {
		s_v, _ := strconv.Atoi(v)

		if s_v == m.activeDate {
			output.WriteString(fmt.Sprintf("[%2s] ", v))
		} else {
			output.WriteString(fmt.Sprintf(" %2s  ", v))
		}

		counter++

		if counter == 7 {
			counter = 0
			output.WriteString("\n")
		}
	}

	//read file
	content, err  := os.ReadFile(filePath)
	if err != nil {
		output.WriteString("\nno entry")
	}else{
		output.WriteString(string(content))

	}

	return output.String()
}

func main() {

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
