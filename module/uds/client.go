package uds

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type UDSClient struct {
	conn net.Conn
}

func Connect() (*UDSClient, error) {
	socketPath := GetSocketPath()
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to daemon: %v", err)
	}

	client := &UDSClient{
		conn: conn,
	}

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err)
				return
			}

			log.Println(message)
		}
	}()

	return client, nil
}
