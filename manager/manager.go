package manager

import (
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
	"log"
	"time"
)

type ClipboardEntry struct {
	Content   string
	Timestamp time.Time
}

type ClipboardManager struct {
	History []ClipboardEntry
}

func NewClipboardManager() *ClipboardManager {
	return &ClipboardManager{
		History: []ClipboardEntry{},
	}
}

func (cm *ClipboardManager) AddEntry(content string) {
	entry := ClipboardEntry{
		Content:   content,
		Timestamp: time.Now(),
	}
	cm.History = append(cm.History, entry)
}

func (cm *ClipboardManager) MonitorClipboard(historyList *widget.List) {
	lastCopied := ""
	for {
		copiedBytes := clipboard.Read(clipboard.FmtText)
		copiedText := string(copiedBytes)

		if copiedText != "" && copiedText != lastCopied {
			lastCopied = copiedText
			cm.AddEntry(copiedText)
			log.Printf("New clipboard entry: %s", copiedText)

			historyList.Refresh()
		}

		time.Sleep(1 * time.Second)
	}
}
