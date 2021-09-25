package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	if commandExists(`git`) {
		fmt.Println("git command exist")
	} else {
		fmt.Println("git command not exist")

	}

	if commandExists(`tig`) {
		fmt.Println("tig command exist")
	} else {
		fmt.Println("tig command not exist")

	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}