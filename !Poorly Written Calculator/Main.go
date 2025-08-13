package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main GUI Formatting
func main() {
	application := app.New()
	window := application.NewWindow("PWC")
	window.Resize(fyne.NewSize(360, 260))

	firstEntry := widget.NewEntry()
	firstEntry.SetPlaceHolder("First number")

	secondEntry := widget.NewEntry()
	secondEntry.SetPlaceHolder("Second number")

	operations := []string{"+", "-", "×", "÷"}
	opSelect := widget.NewSelect(operations, func(string) {})
	opSelect.Selected = "+"

	resultLabel := widget.NewLabel("Result: —")

	calculate := func() {
		a, err1 := strconv.ParseFloat(firstEntry.Text, 64)
		b, err2 := strconv.ParseFloat(secondEntry.Text, 64)
		if err1 != nil || err2 != nil {
			resultLabel.SetText("Result: Bro it's a calculator")
			return
		}

		//Caculator Component
		var result float64
		switch opSelect.Selected {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "×":
			result = a * b
		case "÷":
			if b == 0 {
				resultLabel.SetText("Result: division by zero")
				return
			}
			result = a / b
		default:
			resultLabel.SetText("Result: choose operation")
			return
		}

		resultLabel.SetText(fmt.Sprintf("Result: %g", result))

	}

	calcButton := widget.NewButton("Calculate", calculate)

	//More UI
	content := container.NewVBox(
		widget.NewLabel("Poorly Written Calculator"),
		firstEntry,
		secondEntry,
		opSelect,
		calcButton,
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
