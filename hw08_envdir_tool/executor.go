package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	myEnv := env
	for key, value := range myEnv {
		if value.NeedRemove {
			delete(myEnv, key)
			err := os.Unsetenv(key)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else if strings.Contains(key, "=") {
			fmt.Println("Знак = запрещен для обозначения переменной окружения")
			continue
		}
	}
	args := cmd[1:]
	command := cmd[0]
	envs := os.Environ()

	for key, value := range myEnv {
		envs = append(envs, fmt.Sprintf("%s=%s", key, value.Value))
	}
	Cmd := &exec.Cmd{
		Path: command,
		Args: append([]string{command}, args...),
		Env:  envs,
	}
	if filepath.Base(command) == command {
		if lp, err := exec.LookPath(command); err == nil {
			Cmd.Path = lp
		} else {
			return 1
		}
	}
	if len(cmd) > 0 {
		out, err := Cmd.Output()
		if err != nil {
			fmt.Println(err, "Ошибка при выполнении cmd.Output()")
			return 1
		}
		fmt.Println(string(out))
	} else {
		log.Fatal("Недостаточно аргументов для выполнения программы")
	}
	return
}
