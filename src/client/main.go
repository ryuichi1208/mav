package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"unsafe"
)

type ServerInfo struct {
	serverName string
	ipAddr     string
	port       int
}

type fileMeta struct {
	fileName string
	fileParm string
	user     int
	group    int
	fileSize int
	fileType int
}

var (
	si       ServerInfo
	fm       fileMeta
	filePath string
)

func optParse() {
	flag.StringVar(&si.ipAddr, "ipaddr", "0.0.0.0", "help message for \"s\" option")
	flag.IntVar(&si.port, "port", 8080, "help message for \"i\" option (default 1234)")
	flag.StringVar(&filePath, "f", "./", "help message for \"i\" option (default 1234)")
	flag.Parse()

}

func genIpAddr() string {
	ipaddr := fmt.Sprintf("%s:%d", si.ipAddr, si.port)
	return ipaddr
}

func initFileInfo() {
	fm.fileName = "./fileA"
	fm.fileParm = "0755"
	fm.user = 0
	fm.group = 0
	fm.fileSize = 1024
	fm.fileType = 1
}

func client() {
	initFileInfo()

	conn, err := net.Dial("tcp", genIpAddr())
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer conn.Close()

	meta := fmt.Sprintf("%s,%s,%d,%d,%d,%d\n",
		fm.fileName, fm.fileParm, fm.user, fm.group, fm.fileSize, fm.fileType)
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, []byte(meta))
	if err != nil {
		panic(err)
	}

	readBuf, _ := ioutil.ReadAll(buf)

	conn.Write(readBuf)
	conn.Write([]byte("aaa"))
}

func main() {
	structSize := unsafe.Sizeof(si)
	fmt.Println(structSize, filePath)
	optParse()

	client()
}
