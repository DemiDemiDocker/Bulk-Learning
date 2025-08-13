package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Main application GUI function
func main() {
	application := app.New()
	window := application.NewWindow("Poorly Written Calculator")
	window.Resize(fyne.NewSize(360, 260))

	firstEntry := widget.NewEntry()
	firstEntry.SetPlaceHolder("First number")

	secondEntry := widget.NewEntry()
	secondEntry.SetPlaceHolder("Second number")

	operations := []string{"+", "-", "×", "÷"}
	opSelect := widget.NewSelect(operations, nil)
	opSelect.Selected = "+"

	resultLabel := widget.NewLabel("Result: —")

	//Error handling
	calculate := func() {
		a, err1 := strconv.ParseFloat(firstEntry.Text, 64)
		b, err2 := strconv.ParseFloat(secondEntry.Text, 64)
		if err1 != nil || err2 != nil {
			resultLabel.SetText("Result: Bruh, numbers only")
			return
		}
		//Calculator Functionality
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
	}
	// Calculator Functionality
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

	calcButton := widget.NewButton("Calculate", calculate)

	content := container.NewVBox(
		widget.NewLabel("Poorly Written Calculator"),
		firstEntry,
		secondEntry,
		opSelect,
		calcButton,
		resultLabel)

	//Defyne menu options, will add more later on
	toolsMenu := fyne.NewMenu("Options"),
		fyne.NewMenuItem("Light Mode", nil)
}
