package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dates          map[string][]string
	activeDate     int
	activeMonth    string
	inputfield     textarea.Model
	showinputfield bool
}

func getfileinfo(month string, day int) string {
	// user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// filename
	filename := strconv.Itoa(day) + ".txt"

	// full file path
	return filepath.Join(home, "journal", month, filename)
}

func checkDayExists(allDaysInMonth []string, day int) bool{
	for _, v := range allDaysInMonth {
		if strconv.Itoa(day) == v {
			return true
		}
	}
	return false
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

	inputfield := textarea.New()
	inputfield.Focus()
	inputfield.Placeholder = "how's your day going"
	inputfield.SetWidth(40)
	inputfield.SetHeight(40)

	return model{
		dates:          m,
		activeDate:     day,
		activeMonth:    month.String(),
		inputfield:     inputfield,
		showinputfield: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// when text editor is open all keys should be directed there
	if m.showinputfield {
		var cmd tea.Cmd
		m.inputfield, cmd = m.inputfield.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.showinputfield = false

				// get input content
				content := m.inputfield.Value()

				// write to system
				filePath := getfileinfo(m.activeMonth, m.activeDate)
				os.WriteFile(filePath, []byte(content), 0644)

				// reset text area
				m.inputfield.SetValue("")

				return m, nil
			}
		}

		return m, cmd
	}

	var temp int
	switch msg := msg.(type) {

	case tea.KeyMsg:
		key := msg.String()
		switch key {

		case "left":
			temp = m.activeDate - 1
			if checkDayExists(m.dates[m.activeMonth], temp) {
				m.activeDate = temp
			}

		case "right":
			temp = m.activeDate + 1
			if checkDayExists(m.dates[m.activeMonth], temp) {
				m.activeDate = temp
			}

		case "up":
			temp = m.activeDate - 7 
			if checkDayExists(m.dates[m.activeMonth], temp) {
				m.activeDate = temp
			}

		case "down":
			temp = m.activeDate + 7 
			if checkDayExists(m.dates[m.activeMonth], temp) {
				m.activeDate = temp
			}

		case "a":

			// append previous entry
			filePath := getfileinfo(m.activeMonth, m.activeDate)
			content, err := os.ReadFile(filePath)

			if err != nil {
				// file does not exit
			}else{
				// file exists
				m.inputfield.SetValue(string(content))
			}

			m.showinputfield = true

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil

}

func (m model) View() string {

	// read file content
	filePath := getfileinfo(m.activeMonth, m.activeDate)
	content, err := os.ReadFile(filePath)

	if m.showinputfield {

		return fmt.Sprintf("%s\n\n ", m.inputfield.View())

	} else {

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
		if err != nil {
			output.WriteString("\nno entry")
		} else {
			output.WriteString("\n\n" + string(content))
		}

		return output.String()
	}

}

func main() {

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
