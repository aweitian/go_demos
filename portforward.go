package main

import (
	"log"
	"net"
	"network/network"
)


var (
	appAddr = "192.168.31.9:3306"

	localServerAddr = "127.0.0.1:17741"
)

func main() {
	createConnection()
	log.Println("[已断开]")
}

func createConnection() {
	tcpListener, err := network.CreateTCPListener(localServerAddr)
	if err != nil {
		panic(err)
	}
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		local := connectLocal()
		go network.Join2Conn(tcpConn, local)
	}
}


func connectLocal() *net.TCPConn {
	conn, err := network.CreateTCPConn(appAddr)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}
