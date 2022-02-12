package main

import (
	"fmt"
	"gocontr/config"
	"syscall"
)

func SetupRootfs() {
	if err := syscall.Chroot(config.RootFs); err != nil {
		fmt.Println("Chroot error:", err)
		syscall.Exit(1)
	}
	if err := syscall.Chdir("/"); err != nil {
		fmt.Println("Chdir error:", err)
		syscall.Exit(1)
	}
}
