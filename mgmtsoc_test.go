package mgmtsoc

import (
	"bufio"
	"net"
	"sync"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1) // Adding one for the server goroutine

	config := Config{MgmtHost: "localhost", MgmtPort: 8122}

	// Using a channel to capture errors from the server goroutine
	errCh := make(chan error, 1)

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
		errCh <- err
		t.Errorf("Error occurred: %v", err)
	}

	// Start server in a goroutine
	go func() {
		defer wg.Done()
		err := Start(config, onDataCallback, onErrorCallback)
		errCh <- err
	}()

	time.Sleep(time.Second * 1) // Give server time to start

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

	// Wait for server goroutine to finish
	wg.Wait()

	// Check if any errors were reported from the server goroutine
	select {
	case serverErr := <-errCh:
		if serverErr != nil {
			t.Fatalf("Server error: %v", serverErr)
		}
	default:
		// No error
	}
}

