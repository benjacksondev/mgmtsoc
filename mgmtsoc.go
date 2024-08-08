package mgmtsoc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Config holds the server configuration
type Config struct {
	MgmtHost string
	MgmtPort int
}

// Start initializes and starts the TCP server with the given configuration and callbacks.
func Start(config Config, onDataCallback func(cmd string, args []string, conn net.Conn), onErrorCallback func(err error, conn net.Conn)) bool {
	address := fmt.Sprintf("%s:%d", config.MgmtHost, config.MgmtPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return false
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				onErrorCallback(err, conn)
				continue
			}
			go handleConnection(conn, onDataCallback, onErrorCallback)
		}
	}()

	return true
}

// handleConnection manages the connection for incoming data.
func handleConnection(conn net.Conn, onDataCallback func(cmd string, args []string, conn net.Conn), onErrorCallback func(err error, conn net.Conn)) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			onErrorCallback(err, conn)
			return
		}
		cmdline := strings.Fields(strings.TrimSpace(data))
		if len(cmdline) > 0 {
			cmd := cmdline[0]
			args := cmdline[1:]
			onDataCallback(cmd, args, conn)
		}
	}
}
