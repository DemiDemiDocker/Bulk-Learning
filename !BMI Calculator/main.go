package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main GUI Formatting
func main() {
	application := app.New()
	window := application.NewWindow("BMI Calculator")
	window.Resize(fyne.NewSize(360, 260))

	//More UI
	content := container.NewVBox(
		widget.NewLabel("BMI Calculator"),
	)
	window.SetContent(content)

	//Tool Bar Logic and Options
	var toolsMenu *fyne.Menu
	var aboutMenu *fyne.Menu
	var toggleItem *fyne.MenuItem
	currentVariant := application.Settings().ThemeVariant()
	isDark := currentVariant == theme.VariantDark
	initialLabel := "Dark Mode"
	if isDark {
		initialLabel = "Light Mode"
	}
	toggleItem = fyne.NewMenuItem(initialLabel, func() {
		if isDark {
			application.Settings().SetTheme(theme.LightTheme())
			toggleItem.Label = "Dark Mode"
			isDark = false
		} else {
			application.Settings().SetTheme(theme.DarkTheme())
			toggleItem.Label = "Light Mode"
			isDark = true
		}
		window.SetMainMenu(fyne.NewMainMenu(toolsMenu, aboutMenu))
	})

	toolsMenu = fyne.NewMenu("Options",
		toggleItem,
	)

	aboutMenu = fyne.NewMenu("Demi")

	window.SetMainMenu(fyne.NewMainMenu(toolsMenu, aboutMenu))
	window.ShowAndRun()
}
