package utility

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		return
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
