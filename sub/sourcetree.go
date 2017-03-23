package sourcetree

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CheckArg(args []string) bool {
	count := len(args)
	if count < 2 || count > 3 {
		return false
	}
	return true
}

func CallWinMerge(args []string) {
	prevDir, _ := filepath.Abs(".")
	os.Chdir(fmt.Sprintf("cd %s\n", args[1]))
	command := fmt.Sprintf("git windiff %s %s\n", args[2], args[3])
	fmt.Printf(command)
	exec.Command(command)
	os.Chdir(prevDir)

}
