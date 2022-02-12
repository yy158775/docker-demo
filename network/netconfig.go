package network

import (
	//"github.com/docker/libcontainer/netlink"
	"github.com/milosgajdos/tenus"
	"github.com/vishvananda/netns"
	"gocontr/config"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//var lock sync.Mutex
//var ch chan int

func NetConfig() {

	//默认命名空间
	//dehandle,err := netns.Get()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//桥
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

	pid := os.Getpid()
	//线程就行了

	cmd := exec.Command("/proc/self/exe", "net", strconv.Itoa(pid), linkname)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags:syscall.CLONE_VM | syscall.CLONE_FILES | syscall.CLONE_FS,
	//}
	//lock.Lock()
	log.Println("lock")
	Must(cmd.Start())
	//创建新的命名空间
	newhandle, err := netns.New()
	//主要原因是不能退出，一退出原来那个就没了
	if err != nil {
		log.Fatalln(err)
	}
	netns.Set(newhandle)
	//lock.Unlock()
	//log.Println("wait ch")
	//<- ch
	log.Println("newns finished")
	time.Sleep(time.Second * 4)

	/*
		//回去default
		Must(netns.Set(dehandle))
		//设置命名空间
		Must(netlink.NetworkSetNsFd(pair.NetInterface(), int(nshandle))) //应该还可以

		//回来
		Must(netns.Set(nshandle))
		ethIp,ethIpNet,err := net.ParseCIDR("10.0.41.2/24")
		if err != nil {
			log.Fatalln(err)
		}

		linker, err := tenus.NewLinkFrom(linkname)
		if err != nil {
			log.Fatalln(err)
		}

		Must(linker.SetLinkIp(ethIp, ethIpNet))
		linker.SetLinkUp()
	*/

}

//可不可以展开啊
func Must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func NatService() {

}

func Netnsconfig(nspid string, linkname string) {
	time.Sleep(time.Second * 2)
	log.Println("Netnsconfig")
	linker, err := tenus.NewLinkFrom(linkname)
	if err != nil {
		log.Fatalln(err)
	}

	nsint, err := strconv.Atoi(nspid)
	if err != nil {
		log.Fatalln(err)
	}

	//lock.Lock()

	Must(linker.SetLinkNetNsPid(nsint))
	//lock.Unlock()

	ethIp, ethIpNet, err := net.ParseCIDR(config.IpNet)
	if err != nil {
		log.Fatalln(err)
	}
	Must(linker.SetLinkIp(ethIp, ethIpNet))
	Must(linker.SetLinkUp())

	//log.Println("send ch")
	//ch <-1
	log.Println("set finished")
}
