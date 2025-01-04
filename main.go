package main

import (
	"clipboard/manager"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"golang.design/x/clipboard"
	"log"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Clipboard Manager")

	cm := manager.NewClipboardManager()
	ui, historyList := cm.CreateUI()

	go cm.MonitorClipboard(historyList)

	myWindow.SetContent(ui)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
