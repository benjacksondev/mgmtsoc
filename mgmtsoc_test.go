package mgmtsoc

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	config := Config{MgmtHost: "localhost", MgmtPort: 8122}

	onDataCallback := func(cmd string, args []string, conn net.Conn) {
		expectedCmd := "TEST"
		if cmd != expectedCmd {
			t.Errorf("expected command %s, got %s", expectedCmd, cmd)
		}
		expectedArgs := []string{"arg1", "arg2"}
		for i, arg := range expectedArgs {
			if args[i] != arg {
				t.Errorf("expected arg %s, got %s", arg, args[i])
			}
		}
		conn.Write([]byte("Command received\n"))
	}

	onErrorCallback := func(err error, conn net.Conn) {
		t.Errorf("Error occurred: %v", err)
	}

	go Start(config, onDataCallback, onErrorCallback)
	time.Sleep(time.Second) // Give server time to start

	conn, err := net.Dial("tcp", "localhost:8122")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	message := "TEST arg1 arg2\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatalf("Failed to send data: %v", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	expectedResponse := "Command received\n"
	if response != expectedResponse {
		t.Errorf("expected response %s, got %s", expectedResponse, response)
	}
}
