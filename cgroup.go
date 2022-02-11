package main

import (
	"encoding/binary"
	"os"
	"strings"
)

const (
	DIR_CGROUP_PIDS = "/sys/fs/cgroup/pids/"
	DIR_CGROUP_MEM  = "/sys/fs/cgroup/memory/"
	CGROUP_PROCS    = "/cgroup.procs"
	PIDS_MAX        = "/pids.max"
	MEM_MAX_USGAE   = "/memory.limit_in_bytes"
	MAX_PIDS        = 10
	MAX_MEM         = 1024 * 1024
)

func LimitResource() {
	CreatePidCgroup()
	CreateMemCgroup()
}

func CreatePidCgroup() {
	var builder strings.Builder
	builder.WriteString(DIR_CGROUP_PIDS)
	builder.WriteString(ContainerId)
	path := builder.String()

	os.Mkdir(path, os.ModePerm)

	builder.WriteString(CGROUP_PROCS)
	procfile := builder.String()
	pid := os.Getpid()

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(pid))
	os.WriteFile(procfile, b, os.ModePerm)

	binary.LittleEndian.PutUint32(b, uint32(MAX_PIDS))
	path = path + PIDS_MAX
	os.WriteFile(path, b, os.ModePerm)
}

func CreateMemCgroup() {
	var builder strings.Builder
	builder.WriteString(DIR_CGROUP_MEM)
	builder.WriteString(ContainerId)
	path := builder.String()

	os.Mkdir(path, os.ModePerm)

	builder.WriteString(CGROUP_PROCS)
	procfile := builder.String()
	pid := os.Getpid()

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(pid))
	os.WriteFile(procfile, b, os.ModePerm)

	binary.LittleEndian.PutUint32(b, uint32(MAX_MEM))
	path = path + MEM_MAX_USGAE
	os.WriteFile(path, b, os.ModePerm)
}
