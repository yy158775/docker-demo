package main

import (
	"gocontr/config"
	"gocontr/network"
	"gocontr/volume"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func child() {
	getwd, _ := os.Getwd()
	log.Println("Dir:", getwd)
	//初始化
	config.ConfigRead(os.Args[2])
	//json文件名
	if config.VolumeSrc != "" {
		volume.MountVolume(config.VolumeSrc, config.VolumeDst, config.Mode)
	}
	//NetConfig()
	LimitResource()
	network.NetConfig()
	//JoinNetworkNs(config.NetNs)

	SetupRootfs()
	SetHostname()

	MountProc()

	log.Println("Create exec")
	ExecProcess()
}

func SetHostname() {
	if err := syscall.Sethostname([]byte(config.HostName)); err != nil {
		log.Fatalln("SetHostName Error:", err)
	}
}
