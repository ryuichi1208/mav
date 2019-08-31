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

func client() {
	conn, err := net.Dial("tcp", genIpAddr())
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer conn.Close()

	sendMsg := "Test Messsage.\n"
	conn.Write([]byte(sendMsg))
}

func main() {
	optParse()

	client()
}
