package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Printf("Error: wrong number of arguments")
		fmt.Printf("Usage: %s [bandit id]\n", os.Args[0])
		return
	}
	command := "ssh"
	argument := fmt.Sprintf("ssh://bandit%s@bandit.labs.overthewire.org:2220", os.Args[1])
	cmd := exec.Command(command, argument)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
