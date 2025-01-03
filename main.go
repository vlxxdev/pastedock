package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"time"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	fmt.Println("Clipboard monitor is running in the background...")

	go func() {
		lastCopied := ""
		for {
			copiedBytes := clipboard.Read(clipboard.FmtText)
			copiedText := string(copiedBytes)
			if copiedText != "" && lastCopied != copiedText {
				lastCopied = copiedText
				fmt.Printf("New clipboard content: %s\n", copiedText)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}
