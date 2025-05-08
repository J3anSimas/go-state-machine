# State Machine TUI in Go

This project is a simple implementation of a **finite state machine (FSM)** in Go, with a terminal user interface built using the [Bubbletea](https://github.com/charmbracelet/bubbletea) library.

The main idea is to demonstrate how to use **interfaces** to represent different states and transitions, encapsulating state behavior in a clean and extensible way.

> 🎥 Project inspired by this YouTube video:
> [How to Program in Unity: State Machines Explained](https://www.youtube.com/watch?v=Vt8aZDPzRjI) by **[iHeartGameDev](https://www.youtube.com/@iHeartGameDev)**

## 🧠 Key Concepts

* **Interface-oriented programming**
* **Finite State Machine (FSM) modeling**
* **Terminal-based user interaction using Bubbletea**

## 📦 Technologies Used

* [Go (Golang)](https://golang.org/) `1.24`
* [Bubbletea](https://github.com/charmbracelet/bubbletea) — TUI framework in Go

## 📁 Project Structure

```bash
.
├── main.go               # Application entry point
├── game_state.go         # Game State interface
├── menu_state.go         # Concrete implementation of the menu state
├── game_running_state.go # Concrete implementation of the running game state
├── go.mod
└── README.md
```
## 🚀 Getting Started

1. Clone the repository:

```bash
git clone https://github.com/j3ansimas/state-machine-go.git
cd state-machine-go
```

2. Download dependencies:

```bash
go mod tidy
```

3. Run the project:

```bash
go run main.go
```

## 🔄 How It Works

Each game state implements the `GameState` interface:

```go
type GameState interface {
    EnterState(*Model)
    Update(*Model, msg tea.Msg) (GameState, tea.Cmd)
    View() string
}
```

The main program holds a reference to the current state and delegates the `Update` and `View` methods to it. Each state can return a new state when a transition is required.

## 💡 Future Improvements

* Add a pause state and transition logic
* Support for saving game progress
* Add tests to validate state transitions
* Export FSM logic as a reusable package

## 🤝 Contributions

**Feel free to open issues, suggest improvements, or even submit pull requests! Your feedback and contributions are very welcome.**
