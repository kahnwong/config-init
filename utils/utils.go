package utils

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(filePath string, data string, permission os.FileMode) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Chmod(permission)
	if err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}
}
