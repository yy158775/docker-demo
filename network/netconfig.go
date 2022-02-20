package network

import (
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
		log.Printf("The Network Namespace %s has existed\n", config.NetNs)
	}

	nsHandle, err := netns.GetFromName(config.NetNs)
	if err != nil {
		log.Fatalln(err)
	}

	//获取bridge
	bridge, err := tenus.BridgeFromName(config.Bridge)

	//如果没有test0
	if err != nil {
		err = exec.Command("sudo", "ip", "link", "add", "test0", "type", "bridge").Run()
		if err != nil {
			log.Fatalln(err)
		}
		bridge, err = tenus.BridgeFromName("test0")
		if err != nil {
			log.Fatalln(err)
		}

		//配置IP地址
		ip, ipnet, err := net.ParseCIDR(config.Gateway)
		if err != nil {
			log.Fatalln(err)
		}
		err = bridge.SetLinkIp(ip, ipnet)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err = bridge.SetLinkUp()

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

	log.Println("newns finished")

	//设置命名空间
	Must(netlink.NetworkSetNsFd(pair.NetInterface(), int(nsHandle))) //应该还可以

	err = netns.Set(nsHandle)
	if err != nil {
		log.Fatalln(err)
	}

	ethIp, ethIpNet, err := net.ParseCIDR(config.IpNet)
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
	log.Println(string(output))
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
