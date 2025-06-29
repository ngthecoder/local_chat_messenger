package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {
	sAddr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8090,
		Zone: "",
	}

	con, err := net.DialUDP("udp", nil, sAddr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for true {
		fmt.Printf("Type your message: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cMessage := scanner.Text()

		con.Write([]byte(cMessage))

		if cMessage == "exit" {
			con.Close()
			break
		}

		buff := make([]byte, 1024)
		n, sAddr, err := con.ReadFromUDP(buff)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		sMessage := string(buff[:n])
		if cMessage == "exit" {
			con.Close()
			break
		}

		fmt.Printf("%s(%d): %s\n", sAddr.IP, sAddr.Port, sMessage)
	}

	fmt.Println("Connection Closed!")
}
