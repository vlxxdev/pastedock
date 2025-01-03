package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
	"log"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Clipboard")

	textLabel := widget.NewLabel("Clipboard is empty")

	refreshButton := widget.NewButton("Refresh Clipboard", func() {
		copiedBytes := clipboard.Read(clipboard.FmtText)
		copiedText := string(copiedBytes)
		if copiedText == "" {
			textLabel.SetText("Clipboard is empty")
		} else {
			textLabel.SetText(copiedText)
		}
	})

	clearButton := widget.NewButton("Clear Clipboard", func() {
		clipboard.Write(clipboard.FmtText, []byte(""))
		log.Println("Clipboard has been cleared.")
	})

	content := container.NewVBox(
		textLabel,
		refreshButton,
		clearButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))

	myWindow.ShowAndRun()
}
