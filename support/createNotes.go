package support

import (
	"strings"
	"fmt"
	"os"
)

func CanCreateStorage() (create bool) {
	f, err := os.Open(StoragePath)
	HandleError(err)

	files, err := f.ReadDir(0)
	HandleError(err)

	for _, v := range files {
		
		if v.Name() == StorageDirName && v.IsDir() {
			create = false
			break
		} else {
			create = true
		}
	}
	return
}


func CreateNote(title, content string) (response bool) {
	newNotePath := fmt.Sprintf(StoragePath+"%s/%s_%s_%s.txt", StorageDirName, 
							     strings.Replace(strings.Replace(title, "-", DashReplacement, -1)," ", "-", -1),
								 NoteFileLabelName, GetCurrentTime())

	// Create note file
	f, err := os.Create(newNotePath)

	HandleError(err)
	defer f.Close()

	// Change file mode to RDWR
	err = os.Chmod(newNotePath, 0666)
	HandleError(err)

	// Write the content to the create note
	_, err = f.Write([]byte(content + "\n"))
	HandleError(err)

	response = true
	fmt.Println("Inserted note succesfully")

	return
}

