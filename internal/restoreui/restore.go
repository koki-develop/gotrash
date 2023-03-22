package restoreui

import tea "github.com/charmbracelet/bubbletea"

var (
	_ tea.Model = (*Model)(nil)
)

type Model struct{}

func New() *Model {
	return &Model{}
}

func Run(m *Model) error {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	return "restore ui"
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
