package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"Gantay/weather/internal/api"
	"Gantay/weather/internal/config"
)

type model struct {
	weather  *api.Weather
	settings config.Settings
	help     help.Model
	keys     keyMap
	err      error
}

type keyMap struct {
	quit    key.Binding
	refresh key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view.
// This satisfies the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.quit, k.refresh}
}

// FullHelp returns keybindings for the expanded help view.
// This satisfies the help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.quit, k.refresh}, // Add your key bindings in groups as you want them to appear
	}
}

func defaultKeyMap() keyMap {
	return keyMap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		refresh: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "refresh weather"),
		),
	}
}

func InitialModel(settings config.Settings) model {
	return model{
		settings: settings,
		help:     help.New(),
		keys:     defaultKeyMap(),
	}
}

func (m model) Init() tea.Cmd {
	return fetchWeatherCmd(m.settings)
}

func fetchWeatherCmd(settings config.Settings) tea.Cmd {
	return func() tea.Msg {
		weather, err := api.FetchCurrentWeather(settings.Location, settings.ApiKey)
		if err != nil {
			return errMsg{err}
		}
		return weatherMsg{weather}
	}
}

type weatherMsg struct {
	weather *api.Weather
}

type errMsg struct {
	err error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.refresh):
			return m, fetchWeatherCmd(m.settings)
		}

	case weatherMsg:
		m.weather = msg.weather
		return m, nil

	case errMsg:
		m.err = msg.err
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.weather == nil {
		return "Loading weather data..."
	}

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		Render("Weather Dashboard")

	details := fmt.Sprintf(`
Location: %s, %s
Temperature: %.1f°C
Condition: %s
Wind Speed: %.1f kph
Humidity: %d%%
`,
		m.weather.Location.Name,
		m.weather.Location.Country,
		m.weather.Current.TemC,
		m.weather.Current.Condition.Text,
		m.weather.Current.WindSpeed,
		m.weather.Current.Humidity,
	)

	// The help view now correctly uses the keyMap that implements help.KeyMap
	helpView := m.help.View(m.keys)

	return fmt.Sprintf("%s\n\n%s\n\n%s", title, details, helpView)
}

func Run(settings config.Settings) {
	p := tea.NewProgram(InitialModel(settings))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
