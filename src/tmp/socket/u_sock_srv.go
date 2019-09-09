package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
	os.Remove(path)
	listener, err := net.Listen("unix", path)
	if err != nil {
		return
	}
	defer listener.Close()
	fmt.Println("Server is running" + path)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				return
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				return
			}
			fmt.Println(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello world\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
