package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecProcess() {
	fmt.Println(os.Args[2])
	cmd := exec.Command(os.Args[3], os.Args[4:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ExecProcess Error:", err)
		os.Exit(1)
	}
}
