package public

import (
	"fmt"
	"os"
	"strings"
)

func ParsePublic(pathname string) string {

	lastIdx := strings.LastIndex(pathname, "/")

	if lastIdx == len(pathname)-1 {
		lastIdx = strings.LastIndex(pathname[:len(pathname)-1], "/")
	}

	if lastIdx == -1 {
		return "No Such File"
	}

	fileName := pathname[lastIdx+1 : len(pathname)-1]
	filePath := pathname[:lastIdx]

	constructDir := fmt.Sprintf("./public%s", filePath)
	files, err := os.ReadDir(constructDir)

	if err != nil {
		return "<h4>Error 404 File Not Found</h4>"
	}

	var fileData string
	for _, name := range files {
		fileInfo, _ := name.Info()

		actualName := strings.ToLower(fileInfo.Name())
		filePath := strings.ToLower(fileName)

		if strings.Contains(actualName, filePath) {
			constructDir := fmt.Sprintf("./public%s", strings.Trim(pathname, " "))

			file, readErr := os.ReadFile(constructDir)
			if readErr != nil {
				fileData = string("404 File Not Found")
				break
			}
			fileData = string(file)
			break
		}
		fileData = "404 File Doesnt Exist"
	}

	return fileData

}
