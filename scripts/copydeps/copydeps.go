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

const (
	regExpEveryMinFile   = `^.+\.min\.(js|css)$`
	regExpEveryWoff2File = `^.+\.woff2$`
)

const (
	staticPath = "internal/app/server/static/dist"
)

const (
	jquerySourcePath = "node_modules/jquery/dist"
	jqueryStaticPath = staticPath + "/jquery.com"
)

const (
	fomanticSourcePath = "node_modules/fomantic-ui/dist"
	fomanticStaticPath = staticPath + "/fomantic-ui.com"
)

const (
	lodashSourcePath = "node_modules/lodash"
	lodashStaticPath = staticPath + "/lodash.com"
)

func Panic(who error, what, where string) {
	log.Panicf("(%s): %s: [%s]", where, what, who)
}

func CreateDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		Panic(err, fmt.Sprintf("failed to create \"%s\"", path), "CreateDirectory")
	}

	log.Printf("created: \"%s\"", path)
}

func CopyFile(filename, sourcePath, targetPath string) {
	data, err := os.ReadFile(filepath.Join(sourcePath, filepath.Base(filename)))

	if err != nil {
		Panic(err, fmt.Sprintf("failed to read \"%s\" from \"%s\"", filename, sourcePath), "CopyFile")
	}

	err = os.WriteFile(filepath.Join(targetPath, filepath.Base(filename)), data, os.ModePerm)

	if err != nil {
		Panic(err, fmt.Sprintf("failed to write \"%s\" on \"%s\"", filename, targetPath), "CopyFile")
	}

	log.Printf("created: \"%s\"", filepath.Join(targetPath, filepath.Base(filename)))
}

func CopyDeps(filenames []string, sourcePath, targetPath string) {
	files, err := os.ReadDir(sourcePath)

	if err != nil {
		Panic(err, fmt.Sprintf("failed to copy \"%s\" from \"%s\"", filenames, sourcePath), "CopyDeps")
	}

	CreateDirectory(targetPath)

	if strings.HasPrefix(filenames[0], "^") && strings.HasSuffix(filenames[0], "$") {
		isMinFile := regexp.MustCompile(filenames[0]).MatchString

		for _, file := range files {
			if isMinFile(file.Name()) {
				CopyFile(file.Name(), sourcePath, targetPath)
			}
		}

		return
	}

	for _, filename := range filenames {
		for _, file := range files {
			if filepath.Base(filename) == file.Name() {
				CopyFile(filename, sourcePath, targetPath)
			}
		}
	}
}

func main() {
	err := os.RemoveAll(staticPath)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		Panic(err, fmt.Sprintf("failed to remove \"%s\"", staticPath), "main")
	}

	CreateDirectory(staticPath)

	CopyDeps([]string{"jquery.min.js"}, jquerySourcePath, jqueryStaticPath)

	CopyDeps([]string{"semantic.min.js", "semantic.min.css"}, fomanticSourcePath, fomanticStaticPath)
	CopyDeps([]string{regExpEveryMinFile}, filepath.Join(fomanticSourcePath, "components"), filepath.Join(fomanticStaticPath, "components"))
	CopyDeps([]string{regExpEveryWoff2File}, filepath.Join(fomanticSourcePath, "themes/default/assets/fonts"), filepath.Join(fomanticStaticPath, "themes/default/assets/fonts"))

	CopyDeps([]string{"lodash.min.js"}, lodashSourcePath, lodashStaticPath)
}
