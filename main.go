package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type Model struct {
	currentState     GameState
	menuState        *MenuState
	runningGameState *GameRunningState
}

func initialModel() tea.Model {
	menuState := NewMenuState()
	runningGameState := NewGameRunningState()
	m := Model{
		menuState:        menuState,
		runningGameState: runningGameState,
		currentState:     menuState,
	}
	m.currentState.EnterState(&m)
	return &m
}
func (m *Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, c := m.currentState.UpdateState(m, msg)
	return m, c
}
func (m *Model) View() string {
	s := m.currentState.View()
	return s
}
func (m *Model) SwitchState(state GameState) {
	m.currentState = state
	state.EnterState(m)
}
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
