package server

import (
	"geep/module/types"
	"geep/module/util"
	"net"

	"github.com/mitchellh/mapstructure"
)

func (server *Server) killall(conn net.Conn, message map[string]any) error {
	var killallMessage types.KillAllMessage
	resultMessage := types.KillAllResultMessage{
		Type:    "deleteResult",
		Success: false,
	}

	err := mapstructure.Decode(message, &killallMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	errors := server.pm.KillAll()
	if len(errors) > 0 {
		for _, err := range errors {
			server.mainLogger.Errorln(err)
		}
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	resultMessage.Success = true
	err = util.SendMessage(conn, resultMessage)
	if err != nil {
		return err
	}
	return nil
}
