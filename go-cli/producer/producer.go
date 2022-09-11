package producer

import (
	"fmt"
	"github/nopecho/golang-playground/go-cli/constants"
	"github/nopecho/golang-playground/go-cli/env"
	"log"
	"os"
)

const (
	CREATE_DIR_NAME  string      = ".env"
	CREATE_FILE_NAME string      = ".env"
	FILE_MODE        os.FileMode = 511
)

func CreateEnv() {
	path := createDir()
	filePath := createEnvFile(path)
	WriteEnvFile(filePath)
}

func createEnvFile(dirPath string) string {
	fileName := fmt.Sprintf("%s%c%s", dirPath, os.PathSeparator, CREATE_FILE_NAME)
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()
	return fileName
}

func WriteEnvFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, FILE_MODE)
	checkError(err)

	contents := getEnvContents()
	n, err := file.WriteString(string(contents))
	checkError(err)
	fmt.Printf("env write %d bytes \\n", n)
	defer file.Close()
}

func getEnvContents() []byte {
	originEnv := env.NewEnv()
	originEnv.Init()
	m := originEnv.GetMap()

	b := []byte("")
	for key, val := range m {
		a := fmt.Sprintf("%s%s%s\n", key, constants.EnvSeparator, val)
		b = append(b, a...)
	}
	return b
}

func createDir() string {
	homeDir := getHomeDir()
	createDirName := fmt.Sprintf("%s%c%s", homeDir, os.PathSeparator, CREATE_DIR_NAME)
	err := os.Mkdir(createDirName, FILE_MODE)

	if err != nil && !os.IsExist(err) {
		log.Fatalf("ERROR: %s", err)
		panic(err)
	}
	return createDirName
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	checkError(err)
	return homeDir
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		os.Exit(1)
		panic(err)
	}
}
