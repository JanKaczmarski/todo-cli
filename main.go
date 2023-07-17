package main

import (
	"fmt"
	"os"
	"todo-cli/support"
	"flag"
	"strings"
)



func main() {
	// Define Flags
	var action string
	var title string
	var content string
	var quantity int

	// Parse flags
	flag.StringVar(&action, "action", "show", "A type of action to execute, use <action=show> to show notes and <action=create> to create a TODO note")
	flag.StringVar(&title, "title", "", "A title of your TODO note, only use with <action=create>")
	flag.StringVar(&content, "content", "", "The content of your TODO note, only use with <action=create>")
	flag.IntVar(&quantity, "quantity", 0, "The amount of notes that you want to show, Only use when <action=show>, The possible range is >= 0, when 0 provided all notes will be shown, for quantity=2, 2 notes will be shown")
	
	flag.Parse()

	// Show usage if flags are invalid
	if action == "create" && quantity != 0 {
		flag.Usage()
		os.Exit(1)
	} else if quantity < 0 {
		flag.Usage()
		os.Exit(1)
	} else if action == "show" && title != "" || action == "show" && content != "" {
		fmt.Println("Don't use <actioin=show> with title or content provided")
		flag.Usage()
		os.Exit(1)
	} else if quantity < 0 {
		flag.Usage()
		os.Exit(1)
	}

	if action == "show" {
		for _, titleContentSeparated := range support.ShowNotes(quantity) {
			splittedData := strings.Split(titleContentSeparated, support.TitleContentSeparator)
			fmt.Printf("Title: %s \nContent: %s\n", splittedData[0], splittedData[1])
		}
	} else if action == "create" {
		support.CreateNote(title, content)
	}

}


func init() {
	// Works first run only and when smb removed the Storage Folder
	if support.CanCreateStorage() {
		if err := os.Mkdir(support.StoragePath+support.StorageDirName, os.ModePerm); err != nil {
			support.HandleError(err)
		}
	}
}

