package main

import (
	"log"
	"os/exec"
)

func Panic(who error, where string) {
	log.Println("Upgrade failed!")

	log.Println("Please, check 'Error' or undo changes with: make upgrade-reset")

	log.Panicf("Error: (%s): [%s]", where, who)
}

func Run(name string, args ...string) {
	if err := exec.Command(name, args...).Run(); err != nil {
		Panic(err, "Run")
	}
}

func Lint() {
	log.Println("Linting...")
	Run("make", "lint")
}

func main() {
	log.Println("Starting upgrades...")

	Lint()

	log.Println("Upgrading Go")
	Run("make", "upgrade-go")

	log.Println("Upgrading Node")
	Run("make", "upgrade-node")

	log.Println("Upgrading Tooling")
	Run("make", "install-tooling")

	Lint()

	log.Println("Testing...")
	Run("make", "test-unit")

	log.Println("Committing upgrades")
	Run("git", "add", ".", "--update")
	Run("git", "commit", "-m", "chore(deps): upgrade")

	log.Println("Upgrade completed!")
}
