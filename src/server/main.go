package main

import (
	"flag"
	"fmt"
	"net"
)

type ServerInfo struct {
	serverName string
	ipAddr     string
	port       int
}

var si ServerInfo

func optParse() {
	flag.StringVar(&si.ipAddr, "ipaddr", "0.0.0.0", "help message for \"s\" option")
	flag.IntVar(&si.port, "port", 8080, "help message for \"i\" option (default 1234)")

	flag.Parse()

}

func genIpAddr() string {
	ipaddr := fmt.Sprintf("%s:%d", si.ipAddr, si.port)
	return ipaddr
}

func server() {
	ipaddr := genIpAddr()
	listener, err := net.Listen("tcp", ipaddr)
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Accept error: %s\n", err)
		return
	}

	fmt.Println("Recv Message")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	optParse()

	fmt.Println("Server")
	server()
}
