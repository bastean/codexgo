package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getMods() (mods []string) {
	dirs := []string{"src", "tests"}

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, "go.mod") {
				mods = append(mods, path)
			}

			return nil
		})

		if err != nil {
			panic(err)
		}
	}

	return
}

func getDeps(mods []string) (deps map[string][]string) {
	deps = make(map[string][]string)

	for _, mod := range mods {
		dataBytes, err := os.ReadFile(mod)

		mod, _, _ = strings.Cut(mod, "/go.mod")

		if err != nil {
			panic(err)
		}

		require := string(dataBytes)

		directDepsStart := strings.Index(require, "(") + 1
		directDepsEnd := strings.Index(require, ")")

		indirectDepsStart := strings.LastIndex(require, "(") + 1
		indirectDepsEnd := strings.LastIndex(require, ")")

		rawDeps := []string{}

		rawDeps = append(rawDeps, strings.Split(require[directDepsStart:directDepsEnd], "\n")...)
		rawDeps = append(rawDeps, strings.Split(require[indirectDepsStart:indirectDepsEnd], "\n")...)

		for _, dep := range rawDeps {
			dep = strings.TrimSpace(dep)

			if dep != "" {
				dep = strings.Split(dep, " ")[0]
				deps[mod] = append(deps[mod], dep)
			}
		}
	}

	return
}

func upgradeGo(deps map[string][]string) {
	for path, dep := range deps {
		err := exec.Command("/bin/sh", "-c", "cd "+path+" && "+"go get -u "+strings.Join(dep, " ")).Run()

		if err != nil {
			panic(err)
		}
	}
}

func upgradeNode() {
	if err := exec.Command("make", "upgrade-node").Run(); err != nil {
		panic(err)
	}
}

func runLint() {
	if err := exec.Command("make", "lint-check").Run(); err != nil {
		panic(err)
	}
}

func runTests() {
	if err := exec.Command("make", "compose-test").Run(); err != nil {
		panic(err)
	}
}

func commit() {
	if err := exec.Command("git", "add", ".", "--update").Run(); err != nil {
		panic(err)
	}

	if err := exec.Command("git", "commit", "-m", "fix(deps): upgrade dependencies").Run(); err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Upgrade Failed!")
			log.Println("Please, check 'Error' or undo changes with: make upgrade-reset")
			log.Println("Error:", r)
		}
	}()

	log.Println("Upgrading Go dependencies")

	log.Println("Running Tidy")
	runLint()

	mods := getMods()

	deps := getDeps(mods)

	upgradeGo(deps)

	log.Println("Upgrading Node dependencies")
	upgradeNode()

	log.Println("Running Lint")
	runLint()

	log.Println("Running Tests")
	runTests()

	log.Println("Commit changes")
	commit()

	log.Println("Upgrade Completed!")
}
