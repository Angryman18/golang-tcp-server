package public

import (
	"fmt"
	"os"
	"strings"
)

func ParsePublic(pathname string) *os.File {

	lastIdx := strings.LastIndex(pathname, "/")

	if lastIdx == len(pathname)-1 {
		lastIdx = strings.LastIndex(pathname[:len(pathname)-1], "/")
	}

	if lastIdx == -1 {
		return nil
	}

	fileName := pathname[lastIdx+1 : len(pathname)-1]
	filePath := pathname[:lastIdx]

	constructDir := fmt.Sprintf("./public%s", filePath)
	files, err := os.ReadDir(constructDir)

	if err != nil {
		return nil
	}

	var fileReader *os.File = nil
	for _, name := range files {
		fileInfo, _ := name.Info()

		actualName := strings.ToLower(fileInfo.Name())
		filePath := strings.ToLower(fileName)

		if strings.Contains(actualName, filePath) {
			constructDir := fmt.Sprintf("./public%s", strings.Trim(pathname, " "))

			file, err := os.Open(constructDir)
			if err != nil {
				break
			}
			fileReader = file
			break
		}
	}

	return fileReader

}
