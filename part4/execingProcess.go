package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	//Note that the first argument should be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
