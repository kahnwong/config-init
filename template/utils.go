/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package template

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//go:embed templates/*
var templatesFS embed.FS

func writeFile(filePath string, data string, permission os.FileMode) {
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

func WriteConfig(template string, option string, destFile string) {
	// main
	filename := fmt.Sprintf("%s.%s", option, destFile)

	// -- set dest path --
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, destFile)

	// -- write template --
	content, _ := templatesFS.ReadFile(fmt.Sprintf("templates/%s/%s", template, filename))
	writeFile(destPath, string(content), 0664)
}
