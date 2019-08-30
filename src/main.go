package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	serverFlag bool
	clientFlag bool
)

func optParse() {
	flag.BoolVar(&serverFlag, "server", false, "help message \"server\" option")
	flag.BoolVar(&clientFlag, "client", false, "help message \"client\" option")

	flag.Parse()

	if serverFlag == true && clientFlag == true {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}
}

func genIpAddr() string {
	ipaddr := fmt.Sprintf("%s:%d", "0.0.0.0", 8090)
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

func client() {
	conn, err := net.Dial("tcp", "localhost:8090")
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

	if serverFlag == true {
		fmt.Println("Server")
		server()
	} else if clientFlag == true {
		fmt.Println("Client")
		client()
	} else {
		fmt.Println("nothing", serverFlag)
		os.Exit(1)
	}
}
