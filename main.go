package main

import (
	"fmt"
	"os"
	"bufio"
	"os/user"
	"time"
	"strings"
	"log"
)

func main() {
	//err := ioutil.WriteFile(fmt.Sprintf("/var/lib/todo-cli-data/note_%s.txt", getCurrentTime()), []byte('essa'), 0644)
	//handleError(err)
	currentUser, err := user.Current()
	username := currentUser.Username
	handleError(err)
	
	if err = os.Mkdir(fmt.Sprintf("/home/%s/.todo-cli-data", username), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err2 := os.Create(fmt.Sprintf("/home/%s/.todo-cli-data/note_%s.txt", username,  getCurrentTime()))

	handleError(err2)
	defer f.Close()
	
	_, err3 := f.Write([]byte("essa \n"))
	handleError(err3)
	
	fmt.Println("done")

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

func createNote(title, content string) (response bool) {
	// Create file
	os.Create(fmt.Sprintf("/var/lib/todo-cli-data/note_%s.txt", getCurrentTime()))
	// Write to file

	// Approve thoose changes
	return 
}
