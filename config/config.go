package config

import (
	"fmt"
	"github.com/gookit/config/v2"
	//"github.com/gookit/config/v2/yaml"
	"log"
)

var VolumeFrom, VolumeTo, Mode string
var HostName, RootFs, ContainerId, NetNs string
var IpNet, Gateway, Bridge string

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
	Bridge = config.String("Bridge")

	if config.String("VolumeFrom") != "" {
		VolumeFrom = config.String("VolumeFrom")
		VolumeTo = RootFs + config.String("VolumeTo") //注意拼接
		fmt.Println("VolumeTo:", VolumeTo)
		Mode = config.String("Mode")
	}
}
