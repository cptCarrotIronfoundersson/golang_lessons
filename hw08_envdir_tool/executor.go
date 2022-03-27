package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, value := range env {
		if value.NeedRemove {
			continue
		} else if strings.Contains(key, "=") {
			fmt.Println("Знак = запрещен для обозначения переменной окружения")
			continue
		}
		err := os.Unsetenv(key)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		err = os.Setenv(key, value.Value)
		fmt.Println(key, value.Value)
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
