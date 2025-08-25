package main

import (
	"strconv"
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

	// Add a label to display the result
	resultLabel := widget.NewLabel("")

	// UI switch between metrics; only one form visible at a time
	opSelect := widget.NewSelect(operations, func(value string) {
		// Hide both forms initially
		if metricForm != nil {
			metricForm.Hide()
		}
		if imperialForm != nil {
			imperialForm.Hide()
		}

		switch strings.ToLower(value) {
		case "imperial":
			if imperialForm != nil {
				imperialForm.Show()
			}
			// Try to parse values and calculate BMI
			weight, errW := strconv.ParseFloat(imperialWeight.Text, 64)
			height, errH := strconv.ParseFloat(imperialHeight.Text, 64)
			if errW == nil && errH == nil && height > 0 {
				bmi := (weight / (height * height)) * 703
				resultLabel.SetText("BMI: " + strconv.FormatFloat(bmi, 'f', 2, 64))
			} else {
				resultLabel.SetText("")
			}
		case "metric":
			if metricForm != nil {
				metricForm.Show()
			}
			// Try to parse values and calculate BMI
			weight, errW := strconv.ParseFloat(metricWeight.Text, 64)
			height, errH := strconv.ParseFloat(metricHeight.Text, 64)
			if errW == nil && errH == nil && height > 0 {
				bmi := weight / ((height / 100) * (height / 100))
				resultLabel.SetText("BMI: " + strconv.FormatFloat(bmi, 'f', 2, 64))
			} else {
				resultLabel.SetText("")
			}
		}
	})
	opSelect.PlaceHolder = "Select..."

	// Create calculate button
	calculateButton := widget.NewButton("Calculate BMI", func() {
		// Get current selection to determine which form is active
		currentSelection := opSelect.Selected
		switch strings.ToLower(currentSelection) {
		case "imperial":
			weight, errW := strconv.ParseFloat(imperialWeight.Text, 64)
			height, errH := strconv.ParseFloat(imperialHeight.Text, 64)
			if errW == nil && errH == nil && height > 0 {
				bmi := (weight / (height * height)) * 703
				resultLabel.SetText("BMI: " + strconv.FormatFloat(bmi, 'f', 2, 64))
			} else {
				resultLabel.SetText("Please enter valid weight and height values")
			}
		case "metric":
			weight, errW := strconv.ParseFloat(metricWeight.Text, 64)
			height, errH := strconv.ParseFloat(metricHeight.Text, 64)
			if errW == nil && errH == nil && height > 0 {
				bmi := weight / ((height / 100) * (height / 100))
				resultLabel.SetText("BMI: " + strconv.FormatFloat(bmi, 'f', 2, 64))
			} else {
				resultLabel.SetText("Please enter valid weight and height values")
			}
		default:
			resultLabel.SetText("Please select a measurement system first")
		}
	})

	// Add calculate button to both forms
	imperialForm.Add(calculateButton)
	metricForm.Add(calculateButton)

	// Default selection detection (case-insensitive) and initial visibility
	imperialForm.Hide()
	metricForm.Hide()

	// Build content after forms/select are ready
	content := container.NewVBox(
		widget.NewLabel("Measurement System"),
		opSelect,
		imperialForm,
		metricForm,
		resultLabel,
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
