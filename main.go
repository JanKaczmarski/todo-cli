package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const StorageDirName string = "todo-cli-storage"
const StoragePath string = "/var/lib/"

func main() {

}

func getCurrentTime() (formattedTime string) {
	return time.Now().Format("02.01.2006-15:04:05")
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

func getUserInput() (content, title string) {
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Title: ")
	title, err = reader.ReadString('\n')
	handleError(err)
	title = strings.TrimSpace(title)

	fmt.Print("Content: ")
	content, err = reader.ReadString('\n')
	handleError(err)
	content = strings.TrimSpace(content)
	return
}

func canCreateStorage() (create bool) {
	f, err := os.Open(StoragePath)
	handleError(err)

	files, err := f.ReadDir(0)
	handleError(err)

	for _, v := range files {
		if v.Name() == StorageDirName && v.IsDir() {
			create = false
		} else {
			create = true
		}
	}
	return
}

func createNote(title, content string) (response bool) {
	// If data storage dir doesn't exit create it
	if canCreateStorage() {
		if err := os.Mkdir(StoragePath+StorageDirName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	// Create note file
	f, err2 := os.Create(fmt.Sprintf(StoragePath+"%s/%s_%s.txt", StorageDirName, title, getCurrentTime()))

	handleError(err2)
	defer f.Close()

	// Write the content to the create note
	_, err3 := f.Write([]byte(content))
	handleError(err3)

	response = true
	fmt.Println("Inserted note succesfully")

	return
}
