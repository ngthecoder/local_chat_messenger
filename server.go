package main

import (
	"fmt"
	"net"
)

func server() {
	sAddr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8090,
		Zone: "",
	}

	con, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for true {
		buff := make([]byte, 1024)
		n, cAddr, err := con.ReadFromUDP(buff)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		cMessage := string(buff[:n])
		if cMessage == "exit" {
			con.Close()
			break
		}

		fmt.Printf("%s(%d): %s\n", cAddr.IP, cAddr.Port, cMessage)

		sMessage := cMessage + " (echoing from server)"
		con.WriteToUDP([]byte(sMessage), cAddr)
	}

	fmt.Println("Connection Closed!")
}
