package Volumn

import "syscall"

func MountVolumn(src string, dst string, ro bool) {
	var flag uintptr = 0
	if ro {
		flag |= syscall.MS_RDONLY
	}
	syscall.Mount(src, dst, "", flag, "")
}
