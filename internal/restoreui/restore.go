package restoreui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/gotrash/internal/trash"
)

const headerHeight = 1

var (
	_ tea.Model = (*Model)(nil)
)

type keymap struct {
	Up     key.Binding
	Down   key.Binding
	Cancel key.Binding
}

type Model struct {
	// state
	trashList trash.TrashList
	cursor    int
	cancel    bool

	// component
	input textinput.Model

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

	// keymap
	keymap *keymap
}

func New(ts trash.TrashList) *Model {
	ipt := textinput.New()
	ipt.Placeholder = "Filter"
	ipt.Focus()

	return &Model{
		// state
		trashList: ts,
		cursor:    0,
		cancel:    false,
		// component
		input: ipt,
		// window
		windowYPosition: 0,
		// keymap
		keymap: &keymap{
			Up:     key.NewBinding(key.WithKeys("up", "ctrl+p")),
			Down:   key.NewBinding(key.WithKeys("down", "ctrl+n")),
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
	return fmt.Sprintf("%s\n%s", m.headerView(), m.listView())
}

func (m *Model) headerView() string {
	return m.input.View()
}

func (m *Model) listView() string {
	rows := []string{}

	for i, t := range m.trashList {
		if i < m.windowYPosition {
			continue
		}

		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}

		rows = append(rows, fmt.Sprintf("%s%s", cursor, t.Path))
		if i+1 == m.windowYPosition+(m.windowHeight-headerHeight) {
			break
		}
	}

	return strings.Join(rows, "\n")
}

/*
 * update
 */

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Cancel):
			// cancel
			m.cancel = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Up):
			// up
			m.cursorUp()
		case key.Matches(msg, m.keymap.Down):
			// down
			m.cursorDown()
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.input.Width = m.windowWidth - 3
		m.fixYPosition()
	}

	var cmds []tea.Cmd

	{
		ipt, cmd := m.input.Update(msg)
		m.input = ipt
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) cursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
	m.fixYPosition()
}

func (m *Model) cursorDown() {
	if m.cursor+1 < len(m.trashList) {
		m.cursor++
	}
	m.fixYPosition()
}

func (m *Model) fixYPosition() {
	if m.cursor < m.windowYPosition {
		m.windowYPosition = m.cursor
	}
	if m.cursor+1 >= (m.windowHeight-headerHeight)+m.windowYPosition {
		m.windowYPosition = m.cursor + 1 - (m.windowHeight - headerHeight)
	}
}
