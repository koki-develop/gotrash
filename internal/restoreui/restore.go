package restoreui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
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

	// component
	input textinput.Model

	// keymap
	keymap *keymap
}

func New() *Model {
	ipt := textinput.New()
	ipt.Placeholder = "Filter"
	ipt.Focus()

	return &Model{
		// state
		cancel: false,
		// component
		input: ipt,
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
	return tea.Batch(
		textinput.Blink,
		tea.EnterAltScreen,
	)
}

/*
 * view
 */

func (m *Model) View() string {
	return m.input.View()
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
	case tea.WindowSizeMsg:
		m.input.Width = msg.Width - 3
	}

	var cmds []tea.Cmd

	{
		ipt, cmd := m.input.Update(msg)
		m.input = ipt
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
