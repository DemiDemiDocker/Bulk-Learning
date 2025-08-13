package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Poorly Written Calculator")
	window.Resize(fyne.NewSize(460, 320))

	firstEntry := widget.NewEntry()
	firstEntry.SetPlaceHolder("First number")

	secondEntry := widget.NewEntry()
	secondEntry.SetPlaceHolder("Second number")

	operations := []string{"+", "-", "×", "÷"}
	opSelect := widget.NewSelect(operations, nil)
	opSelect.Selected = "+"

	resultLabel := widget.NewLabel("Result: —")

	calculate := func() {
		a, err1 := strconv.ParseFloat(firstEntry.Text, 64)
		b, err2 := strconv.ParseFloat(secondEntry.Text, 64)
		if err1 != nil || err2 != nil {
			resultLabel.SetText("Result: invalid input")
			return
		}

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

	content := container.NewVBox(
		widget.NewLabel("Poorly Written Calculator"),
		firstEntry,
		secondEntry,
		opSelect,
		calcButton,
		resultLabel,
	)

	// Tools > Open Defyne
	openDefyne := func() {
		cmdName := "defyne"
		if runtime.GOOS == "windows" {
			cmdName = "defyne.exe"
		}
		cmd := exec.Command(cmdName)
		if err := cmd.Start(); err != nil {
			dialog := widget.NewLabel("Defyne not found. Install: go install github.com/fyne-io/defyne@latest")
			window.SetContent(container.NewVBox(content, dialog))
			return
		}
	}

	toolsMenu := fyne.NewMenu("Tools",
		fyne.NewMenuItem("Open Defyne", func() { openDefyne() }),
	)
	window.SetMainMenu(fyne.NewMainMenu(toolsMenu))

	window.SetContent(content)
	window.ShowAndRun()
}
