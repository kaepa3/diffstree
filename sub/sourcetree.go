package sourcetree

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckArg(args []string) bool {
	count := len(args)
	if count < 2 || count > 3 {
		return false
	}
	return true
}

const (
	batfile = "temp.bat"
)

func MakeBatAndDo(args []string) {
	file, err := os.OpenFile(batfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	defd := strings.Replace(args[1], "\\", "\\", -1)

	fmt.Printf(defd)
	file.WriteString(fmt.Sprintf("cd /d %s \n", defd))
	file.WriteString(fmt.Sprintf("git windiff %s %s \n", args[2], args[3]))
}
func GoExec() {
	err := exec.Command(batfile).Start()
	if err != nil {
		panic(err)
	}
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
