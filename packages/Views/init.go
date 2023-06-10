package views

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	clients "github.com/gabrieldebem/clickup/packages/Clients"
	inputs "github.com/gabrieldebem/clickup/packages/Inputs"
	usecases "github.com/gabrieldebem/clickup/packages/UseCases"
)

const SelectListHeight = 14

var (
	SelectTitleStyle      = lipgloss.NewStyle().MarginLeft(2)
	SelectItemStyle       = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	SelectPaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	SelectHelpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	SelectQuitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string { return "" }

type ItemDelegate struct{}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := SelectItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type SelectModel struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m SelectModel) View() string {
	c := clients.ClickUpClient{
		BaseUrl:  os.Getenv("CLICKUP_BASE_URL"),
		Token:    os.Getenv("CLICKUP_TOKEN"),
		SpaceId:  os.Getenv("CLICKUP_SPACE_ID"),
		TeamId:   os.Getenv("CLICKUP_TEAM_ID"),
		FolderId: os.Getenv("CLICKUP_FOLDER_ID"),
		ListId:   os.Getenv("CLICKUP_LIST_ID"),
		UserId:   os.Getenv("CLICKUP_USER_ID"),
	}

	if m.Choice != "" {
		switch m.Choice {
		case "Show only my Tickets":
			return SelectQuitTextStyle.Render(usecases.GetTickets(c, true))
		case "List all Tickets":
			return SelectQuitTextStyle.Render(usecases.GetTickets(c, false))
		case "Show only one Ticket":
			return usecases.ShowTicket(c, "t_2q3q")
		case "Update a ticket":
			ticketData := runUpdateTaskInputs()
			return usecases.UpdateTicket(c, ticketData[0], ticketData[1])
		default:
			return m.Choice
		}
	}
	if m.Quitting {
		return SelectQuitTextStyle.Render("See you later!")
	}

	return "\n" + m.List.View()
}

var program *tea.Program

func RunSelect() {
	items := []list.Item{
		Item("Show only my Tickets"),
		Item("List all Tickets"),
		//Item("Show only one Ticket"),
		Item("Update a ticket"),
	}

	const defaultWidth = 20

	l := list.New(items, ItemDelegate{}, defaultWidth, SelectListHeight)
	l.Title = "What you want to do on ClickUp?"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = SelectTitleStyle
	l.Styles.PaginationStyle = SelectPaginationStyle
	l.Styles.HelpStyle = SelectHelpStyle

	m := SelectModel{List: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// It is running twice, but i dont know why
func runUpdateTaskInputs() []string {
	model := inputs.InitialModel()

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Printf("Could not start program: %s\n", err)
		os.Exit(1)
	}

	var ticketData []string

	for _, input := range model.Inputs {
		ticketData = append(ticketData, input.Value())
	}

	return ticketData
}

