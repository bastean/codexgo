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

const cli = "syncenv"

var envFilesDir string
var envFileModel string

var envFileBackupRegex = regexp.MustCompile(`\.env\..*\.tmp`)

func usage() {
	fmt.Printf("Usage: %v [OPTIONS]\n", cli)
	fmt.Printf("\nE.g.: %v -dir . -model .env.example\n\n", cli)
	flag.PrintDefaults()
}

func backupEnvFiles() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".env") {
			dataBytes, err := os.ReadFile(file.Name())

			if err != nil {
				panic(err)
			}

			err = os.WriteFile(file.Name()+".tmp", dataBytes, 0644)

			if err != nil {
				panic(err)
			}
		}
	}
}

func restoreEnvFilesBackup() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if envFileBackupRegex.Match([]byte(file.Name())) {
			renamed, _ := strings.CutSuffix(file.Name(), ".tmp")
			os.Rename(filepath.Join(envFilesDir, file.Name()), filepath.Join(envFilesDir, renamed))
		}
	}
}

func removeEnvFilesBackup() {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if envFileBackupRegex.Match([]byte(file.Name())) {
			err = os.Remove(file.Name())

			if err != nil {
				panic(err)
			}
		}
	}
}

func getEnvFiles() (envFiles []string) {
	files, err := os.ReadDir(envFilesDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".env") && file.Name() != envFileModel {
			envFiles = append(envFiles, file.Name())
		}
	}

	return
}

func getEnvFileModelVars() []string {
	dataBytes, err := os.ReadFile(envFileModel)

	if err != nil {
		panic(err)
	}

	enVars := strings.Split(string(dataBytes), "\n")

	for i, enVar := range enVars {
		enVars[i] = strings.Split(enVar, "=")[0]
	}

	return enVars
}

func syncEnv(envModelVars []string, envFile string) {
	dataBytes, err := os.ReadFile(envFile)

	if err != nil {
		panic(err)
	}

	envFileVars := strings.Split(string(dataBytes), "\n")

	envFileVarsCleaned := []string{}

	for _, envFileVar := range envFileVars {
		if envFileVar == "" {
			continue
		}

		envFileVarsCleaned = append(envFileVarsCleaned, envFileVar)
	}

	envFileUpdatedVars := ""

	updatedVar := false

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
		panic(err)
	}

	_, err = file.WriteString(envFileUpdatedVars)

	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Sync .env* failed!")

			log.Println("Restoring .env* from backups")
			restoreEnvFilesBackup()

			log.Println("Please, check 'Error' or undo changes with: make sync-env-reset")
			log.Println("Error:", r)
		}
	}()

	flag.StringVar(&envFilesDir, "dir", ".", ".env files directory")
	flag.StringVar(&envFileModel, "model", ".env.example", ".env file model")

	flag.Usage = usage

	flag.Parse()

	log.Println("Creating .env* backups")
	backupEnvFiles()

	log.Println("Searching .env*")
	envFiles := getEnvFiles()
	envFileModelVars := getEnvFileModelVars()

	log.Println("Syncing .env*")
	for _, envFile := range envFiles {
		syncEnv(envFileModelVars, envFile)
	}

	log.Println("Removing .env* backups")
	removeEnvFilesBackup()

	log.Println("Sync .env* completed!")
}
