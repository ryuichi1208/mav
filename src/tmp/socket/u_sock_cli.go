package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
	if err != nil {
		return
	}
	request, err := http.NewRequest("get", "http://localhost:8888", nil)
	if err != nil {
		return
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		return
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return
	}
	fmt.Println(string(dump))
}
