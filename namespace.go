package main

import (
	"github.com/vishvananda/netns"
	"log"
	"strings"
	"syscall"
)

const (
	DIR_NETNS = "/var/run/netns/"
)

func JoinNetworkNs(nsp string) {
	// var name type = value
	// name := value
	builder := strings.Builder{}
	builder.WriteString(DIR_NETNS)
	builder.WriteString(nsp)
	path := builder.String()
	nsHandler, err := netns.GetFromPath(path)
	if err != nil {
		log.Fatalln(nsp + "not found")
	}
	if err := netns.Set(nsHandler); err != nil {
		log.Fatalln("netns Set error:", err)
	}
}

func MountProc() {
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		log.Fatalln("mount proc failed")
	}
}
