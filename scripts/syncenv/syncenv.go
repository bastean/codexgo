package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	cli = "syncenv"
)

var (
	envFilesDir  string
	envFileModel string
)

var (
	envFileBackupRegex = regexp.MustCompile(`^\.env\..*\.tmp$`)
)

func RestoreEnvFilesBackup() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if envFileBackupRegex.MatchString(file.Name()) {
			renamed, _ := strings.CutSuffix(file.Name(), ".tmp")
			os.Rename(filepath.Join(envFilesDir, file.Name()), filepath.Join(envFilesDir, renamed))
		}
	}
}

func Panic(who error, where string) {
	log.Println("Sync .env* failed!")

	log.Println("Restoring .env* from backups")
	RestoreEnvFilesBackup()

	log.Println("Please, check 'Error' or undo changes with: task syncenv-reset")

	log.Panicf("Error: (%s): [%s]", where, who)
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", cli)
	fmt.Printf("\nE.g.: %s -dir . -model .env.example\n\n", cli)
	flag.PrintDefaults()
}

func BackupEnvFiles() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		Panic(err, "BackupEnvFiles")
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".env") {
			data, err := os.ReadFile(file.Name())

			if err != nil {
				Panic(err, "BackupEnvFiles")
			}

			err = os.WriteFile(file.Name()+".tmp", data, 0644)

			if err != nil {
				Panic(err, "BackupEnvFiles")
			}
		}
	}
}

func RemoveEnvFilesBackup() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		Panic(err, "RemoveEnvFilesBackup")
	}

	for _, file := range files {
		if envFileBackupRegex.Match([]byte(file.Name())) {
			err = os.Remove(file.Name())

			if err != nil {
				Panic(err, "RemoveEnvFilesBackup")
			}
		}
	}
}

func GetEnvFiles() (envFiles []string) {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		Panic(err, "GetEnvFiles")
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".env") && file.Name() != envFileModel {
			envFiles = append(envFiles, file.Name())
		}
	}

	return
}

func GetEnvFileModelVars() []string {
	dataBytes, err := os.ReadFile(envFileModel)

	if err != nil {
		Panic(err, "GetEnvFileModelVars")
	}

	enVars := strings.Split(string(dataBytes), "\n")

	for i, enVar := range enVars {
		enVars[i] = strings.Split(enVar, "=")[0]
	}

	return enVars
}

func SyncEnv(envModelVars []string, envFile string) {
	data, err := os.ReadFile(envFile)

	if err != nil {
		Panic(err, "SyncEnv")
	}

	envFileVars := strings.Split(string(data), "\n")

	envFileVarsCleaned := []string{}

	for _, envFileVar := range envFileVars {
		if envFileVar == "" {
			continue
		}

		envFileVarsCleaned = append(envFileVarsCleaned, envFileVar)
	}

	var envFileUpdatedVars string
	var updatedVar bool

	for i, envModelVar := range envModelVars {
		updatedVar = false

		if i+1 == len(envModelVars) {
			break
		}

		if envModelVar == "" {
			envFileUpdatedVars += "\n"
			continue
		}

		for _, envFileVar := range envFileVarsCleaned {
			values := strings.SplitN(envFileVar, "=", 2)
			enVarName := values[0]
			enVarValue := values[1]

			if envModelVar == enVarName {
				envFileUpdatedVars += envModelVar + "=" + enVarValue + "\n"
				updatedVar = true
				break
			}
		}

		if !updatedVar {
			envFileUpdatedVars += envModelVar + "=" + "\n"
		}
	}

	file, err := os.Create(envFile)

	if err != nil {
		Panic(err, "SyncEnv")
	}

	_, err = file.WriteString(envFileUpdatedVars)

	if err != nil {
		Panic(err, "SyncEnv")
	}
}

func main() {
	flag.StringVar(&envFilesDir, "dir", ".", ".env files directory")
	flag.StringVar(&envFileModel, "model", ".env.example", ".env file model")

	flag.Usage = usage

	flag.Parse()

	log.Println("Creating .env* backups")
	BackupEnvFiles()

	log.Println("Searching .env*")
	envFiles := GetEnvFiles()
	envFileModelVars := GetEnvFileModelVars()

	log.Println("Syncing .env*")
	for _, envFile := range envFiles {
		SyncEnv(envFileModelVars, envFile)
	}

	log.Println("Removing .env* backups")
	RemoveEnvFilesBackup()

	log.Println("Sync .env* completed!")
}
