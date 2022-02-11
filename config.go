package main

import (
	"github.com/gookit/config/v2"
	//"github.com/gookit/config/v2/yaml"
	"log"
)

func ConfigRead(filename string) {
	err := config.LoadFiles(filename)
	if err != nil {
		log.Fatalf("Can't find the name:%s", filename)
	}
	HostName = config.String("HostName")
	RootFs = config.String("RootFs")
	ContainerId = config.String("ContainerId")
	NetNs = config.String("NetNs")
}
