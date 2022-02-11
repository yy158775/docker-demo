package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

var ContainerId string
var RootFs string
var HostName string
var NetNs string

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("error")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Parent ERROR:", err)
		os.Exit(1)
	}
}

func child() {
	//初始化
	ConfigRead(os.Args[2])
	//json文件名

	//NetConfig()
	LimitResource()
	JoinNetworkNs(NetNs)

	SetupRootfs()
	SetHostname()

	MountProc()

	ExecProcess()
}

func SetHostname() {
	if err := syscall.Sethostname([]byte(HostName)); err != nil {
		log.Fatalln("SetHostName Error:", err)
	}
}
