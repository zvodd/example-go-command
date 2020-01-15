package util

import (
	"log"
	"net"
)


func FindFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}