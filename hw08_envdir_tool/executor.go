package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, value := range env {
		err := os.Unsetenv(key)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		err = os.Setenv(key, value.Value)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}
	args := cmd[1:]
	command := cmd[0]
	if len(cmd) > 0 {
		command := exec.Command(command, args...)

		out, err := command.Output()
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Println(string(out))
		fmt.Println(err)
	} else {
		log.Fatal("Недостаточно аргументов для выполнения программы")
	}
	return
}
