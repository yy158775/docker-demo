package volume

import (
	"fmt"
	"syscall"
)

func MountVolume(src string, dst string, ro string) {
	var flag uintptr = 0
	if ro != "" {
		flag |= syscall.MS_RDONLY
	}
	if err := syscall.Mount(src, dst, "", flag, ""); err != nil {
		fmt.Printf("Mount %s to %s error:%s", src, dst, err)
	}
}
