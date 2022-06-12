package stoic

import "fmt"

type AboutCommand interface {
	Run()
}

type aboutCommand struct {
	version   string
	buildHash string
}

func NewAboutCommand(version string, buildHash string) AboutCommand {
	return &aboutCommand{
		version:   version,
		buildHash: buildHash,
	}
}

func (a *aboutCommand) Run() {
	about_message := fmt.Sprintf("Version: %s", a.version) + "\n"
	about_message += fmt.Sprintf("Build Hash: %s", a.buildHash) + "\n"

	fmt.Println(about_message)
}
