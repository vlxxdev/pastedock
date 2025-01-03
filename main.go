package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to organize clipboard: %v", err)
	}

	copiedBytes := clipboard.Read(clipboard.FmtText)
	if len(copiedBytes) == 0 {
		fmt.Println("The clipboard is empty or contains non-text data.")
		return
	}

	copiedText := string(copiedBytes)
	fmt.Printf("The clipboard is copied as '%s'.\n", copiedText)
	newText := fmt.Sprintf("Processed text, %s!", copiedText)
	clipboard.Write(clipboard.FmtText, []byte(newText))
	fmt.Printf("The clipboard is copied as '%s'.\n", newText)
}
