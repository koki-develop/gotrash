package restoreui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ tea.Model = (*Model)(nil)
)

type keymap struct {
	Cancel key.Binding
}

type Model struct {
	// state
	cancel bool

	// keymap
	keymap *keymap
}

func New() *Model {
	return &Model{
		// keymap
		keymap: &keymap{
			Cancel: key.NewBinding(key.WithKeys("ctrl+c", "esc")),
		},
	}
}

func Run(m *Model) error {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

/*
 * init
 */

func (m *Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

/*
 * view
 */

func (m *Model) View() string {
	return "hello restore ui"
}

/*
 * update
 */

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Cancel):
			m.cancel = true
			return m, tea.Quit
		}
	}

	return m, nil
}
