package main

import (
	"fmt"
	"game-of-life/pkg/game"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const tickTimeS = 1

var logging *os.File

func main() {
	logfilePath := "./gol.log"
	if logfilePath != "" {
		f, err := tea.LogToFile(logfilePath, "debug")
		if err != nil {
			log.Fatal(err)
		}
		logging = f
	}
	defer logging.Close()

	// TODO: initial state from file
	state := make([][]int, 24)
	for i := range state {
		state[i] = make([]int, 80)
	}

	// start pattern glider
	state[10][11] = 1
	state[11][12] = 1
	state[12][10] = 1
	state[12][11] = 1
	state[12][12] = 1

	p := tea.NewProgram(model{it: 0, maxIt: 1000, state: state, autoplay: false}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	it       int
	maxIt    int
	state    [][]int
	autoplay bool
}

func (m model) Init() tea.Cmd {
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "p":
			m.autoplay = !m.autoplay
			return m, nil
		case "n", "space", " ":
			if !m.autoplay {
				return m.play(nil)
			}
			return m, nil
		}

	case tickMsg:
		if m.autoplay {
			return m.play(tick)
		}
		return m, tick
	}
	return m, nil
}

func (m model) play(tick tea.Cmd) (tea.Model, tea.Cmd) {
	m.state = game.PlayRound(m.state)
	m.it++
	if m.it >= m.maxIt {
		return m, tea.Quit
	}
	return m, tick
}

func (m model) View() string {
	return fmt.Sprintf("Autoplay: %t\nIteration: %d/%d\n\n%s\n", m.autoplay, m.it, m.maxIt, game.ToString(m.state))
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(tickTimeS * time.Second)
	return tickMsg{}
}
