package support

import (
	"time"
	"log"
	"sort"
	"os"
	"strings"
	"fmt"
)

const DateFormat string = "02.01.2006!15:04:05"
const StorageDirName string = "todo-cli-storage"
const StoragePath string = "/var/lib/"
const NoteFileLabelName string = "todo_note_file"
const DashReplacement string = "|+|"
const TitleContentSeparator string = "|sss|"

// Get current time to name newly created note
func GetCurrentTime() (formattedTime string) {
	return time.Now().Format(DateFormat)
}


// Handle Error
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


/// This section is created to extract notes from storage directory

// Sort extracted files by the time of their modification
// func returnes slice of fileNames from latest to oldest modification time
func SortMapByKey(nonEmptyMap map[string]string) (sortedFiles []string){
	keys := make([]string, 0, len(nonEmptyMap))

	for k := range nonEmptyMap{
		keys = append(keys, k)
	}
	sort.Strings(keys)
	

	for _, k := range keys {
		//sortedMap[k] = basket[k]
		sortedFiles = append([]string{nonEmptyMap[k]}, sortedFiles...)
	}

	return 
}


// Extract every file from data-storage and return their names in []string slice
func ListExistingNotes() []string {
	fileModificationTimeMap := map[string]string{}

	f, err := os.Open(StoragePath + StorageDirName)
	HandleError(err)

	files, err := f.ReadDir(0)
	HandleError(err)

	for _, v := range files {
		fileInfo, err := v.Info()
		HandleError(err)
		fileModTime := fileInfo.ModTime()
		
		fileModificationTimeMap[fileModTime.String()] = v.Name()
	}
	sortedFilesSliceByModificationTime := SortMapByKey(fileModificationTimeMap)
	
	return sortedFilesSliceByModificationTime
}


// showNotes-main function, a master of sort to get the output apply 
// users arguments and return prepared data
func ShowNotes(amountToReturn int) []string{
	output := ListExistingNotes()
	outputSlice := []string{}

	if amountToReturn == 0 {
		amountToReturn = len(output)
	}

	if len(output) > amountToReturn {
		output = output[0:amountToReturn]
	} 

	for _, fileName := range output {
		specialSubStringId := strings.Index(fileName, NoteFileLabelName)
		fileContent, err := os.ReadFile(fmt.Sprintf("%s%s/%s",StoragePath, StorageDirName, fileName))
		HandleError(err)
		convertedFileName := strings.Replace(
				strings.Replace(fileName[0:specialSubStringId-1], "-", " ", -1), DashReplacement, "-", -1)
		outputSlice = append([]string{fmt.Sprintf("%s%s%s", convertedFileName, TitleContentSeparator, fileContent)},
								outputSlice...)
	}


	return outputSlice
}


