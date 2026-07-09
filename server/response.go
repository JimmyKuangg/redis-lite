package server

import (
	"net"
)

func WriteResponse(conn net.Conn, response string) error {
	_, err := conn.Write([]byte(response + "\n"))
	return err
}
