package inputs

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	InputFocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	InputBlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	InputCursorStyle         = InputFocusedStyle.Copy()
	InputNoStyle             = lipgloss.NewStyle()
	InputHelpStyle           = InputBlurredStyle.Copy()
	InputCursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	InputFocusedButton = InputFocusedStyle.Copy().Render("[ Submit ]")
	InputBlurredButton = fmt.Sprintf("[ %s ]", InputBlurredStyle.Render("Submit"))
)

var Model InputModel

type InputModel struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

func InitialModel() InputModel {
	m := InputModel{
		Inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = InputCursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Your Ticket Id"
			t.Focus()
			t.PromptStyle = InputFocusedStyle
			t.TextStyle = InputFocusedStyle
		case 1:
			t.Placeholder = "New Ticket Status"
			t.CharLimit = 64
		}

		m.Inputs[i] = t
	}

	Model = m

	return m
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.CursorMode++
			if m.CursorMode > cursor.CursorHide {
				m.CursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := range m.Inputs {
				cmds[i] = m.Inputs[i].Cursor.SetMode(m.CursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.FocusIndex == len(m.Inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.FocusIndex--
			} else {
				m.FocusIndex++
			}

			if m.FocusIndex > len(m.Inputs) {
				m.FocusIndex = 0
			} else if m.FocusIndex < 0 {
				m.FocusIndex = len(m.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := 0; i <= len(m.Inputs)-1; i++ {
				if i == m.FocusIndex {
					// Set focused state
					cmds[i] = m.Inputs[i].Focus()
					m.Inputs[i].PromptStyle = InputFocusedStyle
					m.Inputs[i].TextStyle = InputFocusedStyle
					continue
				}
				// Remove focused state
				m.Inputs[i].Blur()
				m.Inputs[i].PromptStyle = InputNoStyle
				m.Inputs[i].TextStyle = InputNoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *InputModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))

	// Only text Inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m InputModel) View() string {
	var b strings.Builder

	for i := range m.Inputs {
		b.WriteString(m.Inputs[i].View())
		if i < len(m.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &InputBlurredButton
	if m.FocusIndex == len(m.Inputs) {
		button = &InputFocusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(InputHelpStyle.Render("cursor mode is "))
	b.WriteString(InputCursorModeHelpStyle.Render(m.CursorMode.String()))
	b.WriteString(InputHelpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}
