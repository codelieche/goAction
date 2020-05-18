package main

import (
	"fmt"
	"net"
	"os"
)

func parseCIDRDemo() {
	addr := "192.168.1.123/24"
	if ip, ipNet, err := net.ParseCIDR(addr); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println(ip, ipNet)
		fmt.Println(ip.String(), ip.DefaultMask(), ipNet.Mask.String(), ipNet.IP)
	}

	fmt.Println(net.ResolveIPAddr("ip", "www.google.com"))

	fmt.Println(net.LookupPort("127.0.0.1", "3316"))

}

func main() {
	parseCIDRDemo()
}
