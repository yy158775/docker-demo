package network

import (
	"fmt"
	"github.com/docker/libcontainer/netlink"
	"github.com/milosgajdos/tenus"
	"github.com/vishvananda/netns"
	"gocontr/config"
	"log"
	"net"
	"os/exec"
)

func NetConfig() {
	//创建命名空间
	err := exec.Command("sudo", "ip", "netns", "add", config.NetNs).Run()
	if err != nil {
		exec.Command("sudo", "ip", "netns", "delete", config.NetNs).Run()
	}

	//先删除，再增加
	err = exec.Command("sudo", "ip", "netns", "add", config.NetNs).Run()
	if err != nil {
		log.Fatalln(err)
	}

	nsHandle, err := netns.GetFromName(config.NetNs)
	if err != nil {
		log.Fatalln(err)
	}

	bridge, err := tenus.BridgeFromName("test0")
	if err != nil {
		log.Fatalln(err)
	}

	//创建veth pair
	pair, err := tenus.NewVethPair()
	if err != nil {
		log.Fatalln(err)
	}

	Must(bridge.AddSlaveIfc(pair.PeerNetInterface()))

	//保存名字
	linkname := pair.NetInterface().Name

	Must(pair.SetPeerLinkUp())

	Must(bridge.SetLinkUp())

	log.Println("newns finished")

	//设置命名空间
	Must(netlink.NetworkSetNsFd(pair.NetInterface(), int(nsHandle))) //应该还可以

	err = netns.Set(nsHandle)
	if err != nil {
		log.Fatalln(err)
	}

	ethIp, ethIpNet, err := net.ParseCIDR("10.0.41.2/24")
	if err != nil {
		log.Fatalln(err)
	}

	linker, err := tenus.NewLinkFrom(linkname)
	if err != nil {
		log.Fatalln(err)
	}
	Must(linker.SetLinkIp(ethIp, ethIpNet))
	linker.SetLinkUp()

	output, err := exec.Command("sudo", "ip", "route").Output()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(output))

	//设置网关
	//err = exec.Command("sudo","ip", "route", "add", "default", "via", config.Gateway).Run()
	//if err != nil {
	//	log.Fatalln(err)
	//}
}

//可不可以展开啊
func Must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
