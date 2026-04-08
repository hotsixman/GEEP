package client

import (
	"bufio"
	"geep/module/types"
	"geep/module/util"
	"net"
	"strings"
)

func KillAll(conn net.Conn, reader *bufio.Reader) (message *types.KillAllResultMessage, err error) {
	err = util.SendMessage(conn, types.KillAllMessage{Type: "killall"})
	if err != nil {
		return nil, err
	}

	messageJSON, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	message, err = util.ParseMessage[types.KillAllResultMessage]([]byte(strings.TrimSpace(messageJSON)))
	if err != nil {
		return nil, err
	}

	return message, nil
}
