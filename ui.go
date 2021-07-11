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
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
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
		Items: []string{"select_file", "quit"},
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
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
