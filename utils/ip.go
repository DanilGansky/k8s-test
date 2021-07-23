package utils

import (
	"log"
	"net"
	"sync"
)

var (
	ip     string
	onceIP sync.Once
)

func GetOutboundIP() string {
	onceIP.Do(func() {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		ip = conn.LocalAddr().(*net.UDPAddr).String()
	})

	return ip
}
