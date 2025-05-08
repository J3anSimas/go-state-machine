package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuState struct {
	choices []string
	cursor  int
}

func (m *MenuState) EnterState(model *Model) {
	return
}

func (m *MenuState) UpdateState(model *Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return *model, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.choices[m.cursor] {
			case "new game":
				model.SwitchState(model.runningGameState)
			case "join game":
				fmt.Println("join game")
			default:
				return *model, tea.Quit
			}
		}
	}
	return *model, nil
}
func (m *MenuState) View() string {
	s := "Menu\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
func NewMenuState() *MenuState {
	return &MenuState{
		choices: []string{"new game", "join game", "exit game"},
	}
}
