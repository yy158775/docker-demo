package volume

import (
	"fmt"
	"gocontr/config"
	"syscall"
)

func MountVolume(src string, dst string, ro string) {
	var flag uintptr = 0
	//if ro != "" {
	//	flag |= syscall.MS_RDONLY
	//}

	if config.Mode == "RW" {
		flag = syscall.MS_DIRSYNC | syscall.MS_BIND
	}
	if err := syscall.Mount(src, dst, "", flag, ""); err != nil {
		fmt.Printf("Mount %s to %s\nerror:%s\n", src, dst, err)
	}
}
