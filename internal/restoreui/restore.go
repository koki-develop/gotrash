package restoreui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/gotrash/internal/trash"
	"github.com/sahilm/fuzzy"
)

const headerHeight = 1

var (
	_ tea.Model = (*Model)(nil)
)

type keymap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Cancel key.Binding
}

type match struct {
	Trash   *trash.Trash
	Indexes []int
}

type matches []*match

type Model struct {
	// state
	trashList trash.TrashList
	matches   matches
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
			Enter:  key.NewBinding(key.WithKeys("enter")),
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
	var v strings.Builder

	for i, match := range m.matches {
		if i < m.windowYPosition {
			continue
		}

		var s strings.Builder

		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		s.WriteString(cursor)

		for ci, c := range match.Trash.Path {
			style := lipgloss.NewStyle()
			for _, idx := range match.Indexes {
				if ci == idx {
					style = style.Foreground(lipgloss.Color("#00ADD8"))
				}
			}
			if i == m.cursor {
				style = style.Bold(true)
			}
			s.WriteString(style.Render(string(c)))
		}

		v.WriteString(s.String())
		if i+1 == m.windowYPosition+(m.windowHeight-headerHeight) {
			break
		}
		v.WriteString("\n")
	}

	return v.String()
}

/*
 * update
 */

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Enter):
			// enter
			return m, tea.Quit
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
	}

	var cmds []tea.Cmd

	{
		ipt, cmd := m.input.Update(msg)
		m.input = ipt
		cmds = append(cmds, cmd)
	}

	m.filter()
	m.fixYPosition()
	m.fixCursor()

	return m, tea.Batch(cmds...)
}

func (m *Model) cursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *Model) cursorDown() {
	if m.cursor+1 < len(m.matches) {
		m.cursor++
	}
}

func (m *Model) fixCursor() {
	if m.cursor+1 > len(m.matches) {
		m.cursor = len(m.matches) - 1
	}
}

func (m *Model) fixYPosition() {
	if m.windowHeight-headerHeight > len(m.matches) {
		m.windowYPosition = 0
	}
	if m.cursor < m.windowYPosition {
		m.windowYPosition = m.cursor
	}
	if m.cursor+1 >= (m.windowHeight-headerHeight)+m.windowYPosition {
		m.windowYPosition = m.cursor + 1 - (m.windowHeight - headerHeight)
	}
}

func (m *Model) filter() {
	var matches matches

	s := m.input.Value()
	if s == "" {
		for _, t := range m.trashList {
			matches = append(matches, &match{Trash: t})
		}
		m.matches = matches
		return
	}

	fuzzymatches := fuzzy.FindFrom(s, m.trashList)
	for _, fuzzymatch := range fuzzymatches {
		matches = append(matches, &match{
			Trash:   m.trashList[fuzzymatch.Index],
			Indexes: fuzzymatch.MatchedIndexes,
		})
	}
	m.matches = matches
}
