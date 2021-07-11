package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

var clear map[string]func() = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"darwin": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func CreateMainUI() {
	pterm.DefaultCenter.Print(
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(
			pterm.NewStyle(pterm.BgLightBlue),
		).WithMargin(10).Sprint("GoPlay"),
	)
}

func ShowCmdPrompt() string {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"File Selection", "Quit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return ""
	}

	return result
}

func ShowFilePrompt(files []string) string {
	prompt := promptui.Select{
		Label: "Select File or Dir",
		Items: files,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "."
	}

	return result
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		value := clear["linux"]
		value()
	}
}
