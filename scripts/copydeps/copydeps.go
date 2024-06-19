package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const regExpEveryMinFile = `^.+\.min\.(js|css)$`
const regExpEveryWoff2File = `^.+\.woff2$`

const staticPath = "pkg/cmd/server/static/dist"

const jquerySourcePath = "node_modules/jquery/dist"
const jqueryStaticPath = staticPath + "/jquery.com"

const fomanticSourcePath = "node_modules/fomantic-ui/dist"
const fomanticStaticPath = staticPath + "/fomantic-ui.com"

const lodashSourcePath = "node_modules/lodash"
const lodashStaticPath = staticPath + "/lodash.com"

func Panic(msg string, err error) {
	if err != nil {
		log.Panicf("%s: [%s]", msg, err)
	}
}

func createDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		Panic(fmt.Sprintf("Failed to create \"%s\"", path), err)
	}

	log.Printf("Created: \"%s\"", path)
}

func copyFile(filename, sourcePath, targetPath string) {
	data, err := os.ReadFile(filepath.Join(sourcePath, filepath.Base(filename)))

	if err != nil {
		Panic(fmt.Sprintf("Failed to read \"%s\" from \"%s\"", filename, sourcePath), err)
	}

	err = os.WriteFile(filepath.Join(targetPath, filepath.Base(filename)), data, os.ModePerm)

	if err != nil {
		Panic(fmt.Sprintf("Failed to write \"%s\" on \"%s\"", filename, targetPath), err)
	}

	log.Printf("Created: \"%s\"", filepath.Join(targetPath, filepath.Base(filename)))
}

func copyDeps(filenames []string, sourcePath, targetPath string) {
	files, err := os.ReadDir(sourcePath)

	if err != nil {
		Panic(fmt.Sprintf("Failed to copy \"%s\" from \"%s\"", filenames, sourcePath), err)
	}

	createDirectory(targetPath)

	if strings.HasPrefix(filenames[0], "^") && strings.HasSuffix(filenames[0], "$") {
		isMinFile := regexp.MustCompile(filenames[0]).MatchString

		for _, file := range files {
			if isMinFile(file.Name()) {
				copyFile(file.Name(), sourcePath, targetPath)
			}
		}

		return
	}

	for _, filename := range filenames {
		for _, file := range files {
			if filepath.Base(filename) == file.Name() {
				copyFile(filename, sourcePath, targetPath)
			}
		}
	}
}

func main() {
	err := os.RemoveAll(staticPath)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		Panic(fmt.Sprintf("Failed to remove \"%s\"", staticPath), err)
	}

	createDirectory(staticPath)

	copyDeps([]string{"jquery.min.js"}, jquerySourcePath, jqueryStaticPath)

	copyDeps([]string{"semantic.min.js", "semantic.min.css"}, fomanticSourcePath, fomanticStaticPath)
	copyDeps([]string{regExpEveryMinFile}, filepath.Join(fomanticSourcePath, "components"), filepath.Join(fomanticStaticPath, "components"))
	copyDeps([]string{regExpEveryWoff2File}, filepath.Join(fomanticSourcePath, "themes/default/assets/fonts"), filepath.Join(fomanticStaticPath, "themes/default/assets/fonts"))

	copyDeps([]string{"lodash.min.js"}, lodashSourcePath, lodashStaticPath)
}
