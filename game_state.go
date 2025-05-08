package main

import tea "github.com/charmbracelet/bubbletea"

type GameState interface {
	EnterState(*Model)
	UpdateState(*Model, tea.Msg) (Model, tea.Cmd)
	View() string
}
