package manager

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

const maxLength = 50
const removeButtonSymbol = "X"

func (cm *ClipboardManager) CreateUI(mainWindow fyne.Window) (fyne.CanvasObject, *widget.List) {
	var historyList *widget.List

	historyList = widget.NewList(
		func() int {
			return len(cm.History)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
			deleteButton := widget.NewButton(removeButtonSymbol, nil)

			return container.NewHBox(
				container.New(layout.NewMaxLayout(), label),
				container.New(layout.NewCenterLayout(), deleteButton),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			entry := cm.History[i]
			content := entry.Content
			if len(content) > maxLength {
				content = content[:maxLength] + "..."
			}

			hBox := o.(*fyne.Container)
			label := hBox.Objects[0].(*fyne.Container).Objects[0].(*widget.Label)
			deleteButton := hBox.Objects[1].(*fyne.Container).Objects[0].(*widget.Button)

			label.SetText(fmt.Sprintf("[%s] %s", entry.Timestamp.Format("2006-01-02 15:04"), content))
			deleteButton.OnTapped = func() {
				cm.History = append(cm.History[:i], cm.History[i+1:]...)
				historyList.Refresh()
			}
		},
	)

	historyList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(cm.History) {
			entry := cm.History[id]
			clipboard.Write(clipboard.FmtText, []byte(entry.Content))
		}
	}

	clearClipboardButton := widget.NewButton("Clear Clipboard", func() {
		clipboard.Write(clipboard.FmtText, []byte(""))
	})

	clearHistoryButton := widget.NewButton("Clear History", func() {
		cm.History = nil
		historyList.Refresh()
	})

	buttons := container.NewHBox(
		clearClipboardButton,
		clearHistoryButton,
	)

	historyListContainer := container.NewScroll(historyList)

	ui := container.NewBorder(
		nil,
		buttons,
		nil,
		nil,
		historyListContainer,
	)
	return ui, historyList
}
