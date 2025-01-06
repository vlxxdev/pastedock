package manager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Settings struct {
	App fyne.App
}

func NewSettings(app fyne.App) *Settings {
	return &Settings{
		App: app,
	}
}

func (sm *Settings) ShowSettingsWindow() {
	settingsWindow := sm.App.NewWindow("Settings")
	settingsWindow.Resize(fyne.NewSize(350, 500))

	themeOptions := []string{"Light", "Dark", "Auto"}
	themeSelect := widget.NewSelect(themeOptions, func(selected string) {
		switch selected {
		case "Light":
			sm.App.Settings().SetTheme(theme.LightTheme())
		case "Dark":
			sm.App.Settings().SetTheme(theme.DarkTheme())
		case "Auto":
			if sm.isNight() {
				sm.App.Settings().SetTheme(theme.DarkTheme())
			} else {
				sm.App.Settings().SetTheme(theme.LightTheme())
			}
		}
	})
	themeSelect.PlaceHolder = "Select Theme"

	closeButton := widget.NewButton("Close", func() {
		settingsWindow.Close()
	})

	settingsWindow.SetContent(container.NewVBox(
		themeSelect,
		closeButton,
	))

	settingsWindow.Show()
}

func (sm *Settings) isNight() bool {
	hour := time.Now().Hour()
	return hour >= 18 || hour < 6
}
