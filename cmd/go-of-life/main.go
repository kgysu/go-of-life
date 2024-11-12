package main

import (
	"fmt"
	"game-of-life/pkg/game"
	"log"
	"log/slog"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const tickTimeS = 1

var logging *os.File

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	logfilePath := "./gol.log"
	if logfilePath != "" {
		f, err := tea.LogToFile(logfilePath, "debug")
		if err != nil {
			log.Fatal(err)
		}
		logging = f
	}
	defer logging.Close()

	state := make([][]int, 24)
	for i := range state {
		state[i] = make([]int, 80)
	}

	// start pattern
	// - - x x - -
	// - - x x - -
	// x x x - - -
	// - x - - - -
	// state[12][40] = 1
	// state[12][41] = 1
	// state[13][40] = 1
	// state[13][41] = 1
	// state[14][38] = 1
	// state[14][39] = 1
	// state[14][40] = 1
	// state[15][39] = 1

	// state[11][20] = 1
	// state[11][21] = 1
	// state[12][22] = 1
	// state[12][23] = 1
	// state[13][22] = 1
	// state[13][23] = 1
	// state[14][20] = 1
	// state[14][21] = 1

	// state[11][25] = 1
	// state[11][26] = 1
	// state[12][27] = 1
	// state[12][28] = 1
	// state[13][27] = 1
	// state[13][28] = 1
	// state[14][25] = 1
	// state[14][26] = 1

	state[10][10] = 1
	state[10][11] = 1
	state[10][12] = 1
	state[11][11] = 1
	state[11][12] = 1
	state[12][10] = 1
	state[12][12] = 1

	// Initialize our program
	p := tea.NewProgram(model{it: 1000, state: state}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	it    int
	state [][]int
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "space":
			slog.Debug("space pressed")
			log.Println("println space pressed")
			fmt.Println("fmt print space pressed")
			m.state = game.PlayRound(m.state)
			return m, nil
		}

	case tickMsg:
		m.it--
		if m.it <= 0 {
			return m, tea.Quit
		}

		// nextRound
		// m.res = fmt.Sprintf("X XX X %d", m.it)
		// m.state = game.PlayRound(m.state)

		return m, tick
	}
	return m, nil
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	// testString := "Hello!"
	// for i := 0; i < 2500; i++ {
	// 	testString += "#"
	// }
	// testString += "\n"
	// return testString + game.ToString(m.state)
	return fmt.Sprintf("Hi. This program will exit in %d seconds.\n\n%s\n", m.it, game.ToString(m.state))
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(tickTimeS * time.Second)
	return tickMsg{}
}
