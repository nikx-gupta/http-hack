package main

import "fmt"

func main() {
	var msg [512]byte
	msg[0] = 255
	msg[1] = msg[0] << 1
	fmt.Println(msg[0])
	fmt.Println(msg[0] >> 2)
}
