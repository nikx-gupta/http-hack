package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	//service := os.Args[1]
	//addr, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	addr, err := net.ResolveTCPAddr("tcp4", "mail.google.com:443")
	checkError(err)

	fmt.Printf("Resolved: %s\n", addr)

	tcpConn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)

	tcpConn.Write([]byte("CONNECT / HTTP/1.1\r\nHOST: mail.google.com\r\n\r\n"))

	result, err := ioutil.ReadAll(tcpConn)
	checkError(err)

	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		//	fmt.Printf("error resolving: %s", err.Error())
		panic(err)
	}
}
