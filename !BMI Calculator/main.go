package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main GUI Window Formatting
func main() {
	application := app.New()
	window := application.NewWindow("BMI Calculator")
	window.Resize(fyne.NewSize(360, 460))

	//Definitions
	operations := []string{"Imperial", "Metric"}
	var imperialForm *fyne.Container
	var metricForm *fyne.Container

	// Build input forms first
	imperialWeight := widget.NewEntry()
	imperialWeight.SetPlaceHolder("Weight (lbs)")
	imperialHeight := widget.NewEntry()
	imperialHeight.SetPlaceHolder("Height (in)")
	imperialForm = container.NewVBox(
		imperialWeight,
		imperialHeight,
	)

	metricWeight := widget.NewEntry()
	metricWeight.SetPlaceHolder("Weight (kg)")
	metricHeight := widget.NewEntry()
	metricHeight.SetPlaceHolder("Height (cm)")
	metricForm = container.NewVBox(
		metricWeight,
		metricHeight,
	)

	// UI switch between metrics; only one form visible at a time
	opSelect := widget.NewSelect(operations, func(value string) {
		switch strings.ToLower(value) {
		case "imperial":
			if metricForm != nil {
				metricForm.Hide()
			}
			if imperialForm != nil {
				imperialForm.Show()
			}
		case "metric":
			if imperialForm != nil {
				imperialForm.Hide()
			}
			if metricForm != nil {
				metricForm.Show()
			}
		}
	})

	// Default selection detection (case-insensitive) and initial visibility
	opSelect.Selected = "Imperial"
	if strings.ToLower(opSelect.Selected) == "imperial" {
		metricForm.Hide()
		imperialForm.Show()
	} else {
		imperialForm.Hide()
		metricForm.Show()
	}

	// Build content after forms/select are ready
	content := container.NewVBox(
		widget.NewLabel("Measurement System"),
		opSelect,
		imperialForm,
		metricForm,
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
