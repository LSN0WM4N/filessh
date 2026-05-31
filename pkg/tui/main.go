package tui

// These imports will be used later in the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

type model struct {
	cursor int

	explorer *explorer.ExplorerPlugin
	eventBus *bus.EventBus
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.explorer.Entries())-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			newDir := m.explorer.Cwd() + "/" + m.explorer.Entries()[m.cursor].Name() // Wat da HELL?

			m.eventBus.Publish(bus.Event{
				Type:    explorer.EventEnterDir,
				Payload: newDir,
			})
		case "backspace":
			m.eventBus.Publish(bus.Event{
				Type: explorer.EventGoBack,
			})
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	// The header
	var s strings.Builder
	s.WriteString("What should we buy at the market?\n\n")

	// Iterate over our choices
	for i, entry := range m.explorer.Entries() {
		if i > 20 {
			continue
		}

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		fmt.Fprintf(&s, "%s %s\n", cursor, entry)
	}

	// The footer
	s.WriteString("\nPress q to quit.\n")

	// Send the UI for rendering
	return tea.NewView(s.String())
}

func InitTui(explorer *explorer.ExplorerPlugin, eventBus *bus.EventBus) {
	enableDebug()

	p := tea.NewProgram(model{
		eventBus: eventBus,
		explorer: explorer,
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func enableDebug() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}
