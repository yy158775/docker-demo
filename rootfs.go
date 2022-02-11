package main

import (
	"fmt"
	"syscall"
)

func SetupRootfs() {
	if err := syscall.Chroot(RootFs); err != nil {
		fmt.Println("Chroot error:", err)
		syscall.Exit(1)
	}
	if err := syscall.Chdir("/"); err != nil {
		fmt.Println("Chdir error:", err)
		syscall.Exit(1)
	}
}
