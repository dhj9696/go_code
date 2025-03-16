package main

import (
	"fmt"
	"net"
)

func main() {

	//To4将IPv4地址ip转换为4字节表示。如果ip不是IPv4地址，则To4返回nil。
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)
	fmt.Println(ipv6.To4())
	fmt.Println(ipv4.To4())

	//To16将IP地址ip转换为16字节表示。如果ip不是IP地址（长度错误），To16返回nil。
	//ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//ipv4 := net.IPv4(10, 255, 0, 0)
	//fmt.Println(ipv6.To16())
	//fmt.Println(ipv4.To16())

	//IsPrivate根据RFC 1918（IPv4地址）和RFC 4193（IPv6地址）报告ip是否为私有地址。
	//ipv6Private := net.ParseIP("fc00::")
	//ipv6Public := net.ParseIP("fe00::")
	//ipv4Private := net.ParseIP("10.255.0.0")
	//ipv4Public := net.ParseIP("11.0.0.0")
	//fmt.Println(ipv6Private.IsPrivate())
	//fmt.Println(ipv6Public.IsPrivate())
	//fmt.Println(ipv4Private.IsPrivate())
	//fmt.Println(ipv4Public.IsPrivate())

	//IsMulticast报告ip是否是多播地址。
	/*
		多播地址
		定义
		多播（Multicast）是一种一对多的网络通信模式，多播地址是用于标识一组设备的地址。发送方使用多播地址发送数据时，数据会被同时传输到加入了该多播组的所有接收方设备。
		特点
		一对多通信：可以同时将数据发送给多个接收者，减少了数据的重复传输，提高了网络效率。
		资源利用率高：相比于多个单播通信，多播可以在网络中共享数据传输路径，节省了网络带宽和服务器资源。
		动态成员管理：接收者可以动态地加入或离开多播组，使得多播通信更加灵活。
		应用场景
		视频会议：在视频会议中，会议发起者可以将音视频数据通过多播地址发送给所有参会者，避免了为每个参会者单独发送数据，减轻了网络负担。
		网络直播：直播服务器可以使用多播地址将直播流同时发送给大量的观众，提高了直播的效率。
		地址范围
		IPv4：IPv4 多播地址的范围是从 224.0.0.0 到 239.255.255.255。例如，224.0.0.1 表示本地网络中的所有主机，224.0.0.2 表示本地网络中的所有路由器。
		IPv6：IPv6 多播地址以 ff00::/8 开头，例如 ff02::1 表示链路本地范围内的所有节点。
	*/
	//ipv6Multi := net.ParseIP("FF00::")
	//ipv6LinkLocalMulti := net.ParseIP("ff02::1")
	//ipv6Lo := net.ParseIP("::1")
	//ipv4Multi := net.ParseIP("239.0.0.0")
	//ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	//ipv4Lo := net.ParseIP("127.0.0.0")
	//fmt.Println(ipv6Multi.IsMulticast())
	//fmt.Println(ipv6LinkLocalMulti.IsMulticast())
	//fmt.Println(ipv6Lo.IsMulticast())
	//fmt.Println(ipv4Multi.IsMulticast())
	//fmt.Println(ipv4LinkLocalMulti.IsMulticast())
	//fmt.Println(ipv4Lo.IsMulticast())
}
