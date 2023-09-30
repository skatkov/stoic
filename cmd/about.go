package stoic

import "fmt"

type AboutCommand interface {
	Run()
}

type aboutCommand struct {
	version    string
	commitHash string
	date       string
}

func NewAboutCommand(version, commitHash, date string) AboutCommand {
	return &aboutCommand{
		version:    version,
		commitHash: commitHash,
		date:       date,
	}
}

func (a *aboutCommand) Run() {
	about_message := fmt.Sprintf("Version: %s", a.version) + "\n"
	about_message += fmt.Sprintf("Commit Hash: %s", a.commitHash) + "\n"
	about_message += fmt.Sprintf("Build Hash: %s", a.date) + "\n"

	fmt.Println(about_message)
}
