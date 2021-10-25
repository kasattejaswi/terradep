package configs

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func InitializeConfig() {
	usersDirectory, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to get user's home directory: ", err)
	}
	toCreate := filepath.Join(usersDirectory, ".terradep")
	if CheckFileExistence(filepath.Join(toCreate, "repositoryConfig.yml")) {
		log.Fatal(`Configuration is already present. 
To force reinitilization, use -force flag.
This will create a backup of existing configuration under the same directory.`)
	} else {
		WriteConfiguration(toCreate)
	}
}

func ForceInitializeConfig() {
	usersDirectory, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to get user's home directory: ", err)
	}
	toCreate := filepath.Join(usersDirectory, ".terradep")
	CreateFileBackup("repositoryConfig.yml", toCreate)
	WriteConfiguration(toCreate)
}

func CreateFileBackup(fileName string, fileLocation string) {
	newFileName := fileName + "." + strconv.Itoa(int(time.Now().Unix()))
	buff, err := ioutil.ReadFile(filepath.Join(fileLocation, fileName))
	if err != nil {
		log.Fatal("Internal error occurred. Unable to read existing configuration file: ", err)
	}
	err = os.WriteFile(filepath.Join(fileLocation, newFileName), buff, 0755)
	if err != nil {
		log.Fatal("Unable to create backup of existing configurtion file: ", err)
	}
	log.Println("Backup completed successfully: ", newFileName)
}

func WriteConfiguration(pathToFolder string) {
	err := os.MkdirAll(pathToFolder, 0755)
	if err != nil {
		log.Fatal("Failed to create directory for configuration: ", err)
	}
	buff, err := ioutil.ReadFile("configs/base.yml")
	if err != nil {
		log.Fatal("Internal error occurred. Unable to read base.yml: ", err)
	}
	err = os.WriteFile(filepath.Join(pathToFolder, "repositoryConfig.yml"), buff, 0755)
	if err != nil {
		log.Fatal("Unable to create base configuration file: ", err)
	}
	log.Println("Configuration generated successfully")
}

func CheckFileExistence(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		log.Fatal("Internal error. Unable to detect existence of configuration: ", err)
	}
	return false
}
