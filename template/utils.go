/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package template

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func CreateDir(dir string) error {
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, dir)

	err := os.MkdirAll(filepath.Join(destPath), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func WriteConfig(template string, filename string, destFile string) {
	// set dest path
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, destFile)

	// write template
	content, _ := templatesFS.ReadFile(fmt.Sprintf("templates/%s/%s", template, filename))
	writeFile(destPath, string(content), 0664)
}

func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
