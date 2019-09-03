package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type ServerInfo struct {
	serverName string
	ipAddr     string
	port       int
}

var si ServerInfo

func getTime() time.Time {
	return time.Now()
}

func optParse() {
	flag.StringVar(&si.ipAddr, "ipaddr", "0.0.0.0", "help message for \"s\" option")
	flag.IntVar(&si.port, "port", 8080, "help message for \"i\" option (default 1234)")

	t := getTime()
	fmt.Println("[INFO] DATE :", t)
	fmt.Println("[INFO] Accept :", si.ipAddr)
	fmt.Println("[INFO] Port :", si.port)

	flag.Parse()

}

func genIpAddr() string {
	ipaddr := fmt.Sprintf("%s:%d", si.ipAddr, si.port)
	return ipaddr
}

func openFile(filepath string) {
	_, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
}

func fileOperationIntr(metaFileInfo string, cnt int) {
	if cnt == 0 {
		fileBuf := strings.Split(metaFileInfo, ",")
		fmt.Println("[INFO]", metaFileInfo)
		openFile(fileBuf[0])
	}
	fmt.Println(metaFileInfo)
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
	typeBuf := make([]byte, 1024)
	cnt := 0
	for {
		n, err := conn.Read(typeBuf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
		}
		fileOperationIntr(string(typeBuf), cnt)
		cnt++
	}
}

func main() {
	// オプション解析
	optParse()

	// サービス起動処理のエントリポイント
	server()
}
