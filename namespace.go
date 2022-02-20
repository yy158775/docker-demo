package main

import (
	//"github.com/vishvananda/netns"
	"log"
	"strings"
	"syscall"
)

const (
	DIR_NETNS = "/var/run/netns/"
)

func JoinNetworkNs(nsp string) {
	builder := strings.Builder{}
	builder.WriteString(DIR_NETNS)
	builder.WriteString(nsp)
}

func MountProc() {
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		log.Fatalln("mount proc failed")
	}
}
