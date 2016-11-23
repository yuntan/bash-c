package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	args := make([]string, len(os.Args))
	copy(args, os.Args)
	args[0] = filepath.Base(args[0])
	args[0] = strings.TrimSuffix(args[0], ".exe")
	for i, arg := range args {
		args[i] = strings.Replace(arg, " ", `\ `, -1)
	}

	cmd := exec.Command("bash", "-c", strings.Join(args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run() // blocking

	if err == nil {
		return
	}

	if e, ok := err.(*exec.ExitError); ok {
		code := e.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		os.Exit(code)
	} else if err == exec.ErrNotFound {
		fmt.Fprintln(os.Stderr, "bash command not found")
		os.Exit(127)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
