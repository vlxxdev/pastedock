package manager

import (
	"golang.design/x/clipboard"
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

func (cm *ClipboardManager) ClearHistory() {
	cm.History = nil
}

func (cm *ClipboardManager) MonitorClipboard(newEntries chan<- ClipboardEntry) {
	lastCopied := ""
	for {
		copiedBytes := clipboard.Read(clipboard.FmtText)
		copiedText := string(copiedBytes)

		if copiedText != "" && copiedText != lastCopied {
			lastCopied = copiedText
			cm.AddEntry(copiedText)

			newEntries <- ClipboardEntry{
				Content:   copiedText,
				Timestamp: time.Now(),
			}
		}

		time.Sleep(1 * time.Second)
	}
}
