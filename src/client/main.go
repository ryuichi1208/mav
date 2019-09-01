package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
	fileSize int64
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
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	// ファイルの情報を取得する
	fileInfo, er := file.Stat()
	if er != nil {
		return
	}

	fm.fileName = file.Name()
	fm.fileParm = "0755"
	fm.user = 0
	fm.group = 0
	fm.fileSize = fileInfo.Size()
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
	// オプション解析
	optParse()

	// ファイル操作のエントリポイント
	client()
}
