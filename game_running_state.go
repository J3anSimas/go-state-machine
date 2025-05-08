package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GameRunningState struct {
}

func (g *GameRunningState) EnterState(model *Model) {
	return
}

func (g *GameRunningState) UpdateState(model *Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return *model, tea.Quit
		}
	}
	return *model, nil
}
func (g *GameRunningState) View() string {
	s := "Game Running\n\n"

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func NewGameRunningState() *GameRunningState {
	return &GameRunningState{}
}
