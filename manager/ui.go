package manager

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
	"log"
)

func (cm *ClipboardManager) CreateUI() (fyne.CanvasObject, *widget.List) {
	historyList := widget.NewList(
		func() int {
			return len(cm.History)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			entry := cm.History[i]
			o.(*widget.Label).SetText(fmt.Sprintf("[%s] %s", entry.Timestamp.Format("2006-01-02 15:04"), entry.Content))
		},
	)

	historyList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(cm.History) {
			entry := cm.History[id]
			clipboard.Write(clipboard.FmtText, []byte(entry.Content))
			log.Printf("Copied back to clipboard: %s\n", entry.Content)
		}
	}

	clearClipboardButton := widget.NewButton("Clear Clipboard", func() {
		clipboard.Write(clipboard.FmtText, []byte(""))
		log.Println("Clipboard cleared")
	})

	clearHistoryButton := widget.NewButton("Clear History", func() {
		cm.History = nil
		historyList.Refresh()
		log.Println("Clipboard history cleared")
	})

	buttons := container.NewHBox(
		clearClipboardButton,
		clearHistoryButton,
	)

	return container.NewBorder(
		nil,
		buttons,
		nil,
		nil,
		historyList,
	), historyList
}
