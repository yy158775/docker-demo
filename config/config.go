package config

import (
	"github.com/gookit/config/v2"
	//"github.com/gookit/config/v2/yaml"
	"log"
)

var VolumeSrc, VolumeDst, Mode string
var HostName, RootFs, ContainerId, NetNs string
var IpNet, Gateway string

func ConfigRead(filename string) {
	err := config.LoadFiles(filename)
	if err != nil {
		log.Fatalf("Can't find the name:%s", filename)
	}
	HostName = config.String("HostName")
	RootFs = config.String("RootFs")
	ContainerId = config.String("ContainerId")
	NetNs = config.String("NetNs")
	IpNet = config.String("IpNet")
	Gateway = config.String("Gateway")

	if config.String("VolumeSrc") != "" {
		VolumeSrc = config.String("VolumeSrc")
		VolumeDst = config.String("VolumeDst")
		Mode = config.String("Mode")
	}
}
