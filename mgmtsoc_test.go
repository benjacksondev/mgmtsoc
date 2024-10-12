package mgmtsoc

import (
  "fmt"
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

	// Using a channel to capture data callback errors
	dataCallbackErrCh := make(chan error, 1)

	onDataCallback := func(cmd string, args []string, conn net.Conn) {
		expectedCmd := "TEST"
		if cmd != expectedCmd {
			dataCallbackErrCh <- fmt.Errorf("expected command %s, got %s", expectedCmd, cmd)
			return
		}
		expectedArgs := []string{"arg1", "arg2"}
		for i, arg := range expectedArgs {
			if args[i] != arg {
				dataCallbackErrCh <- fmt.Errorf("expected arg %s, got %s", arg, args[i])
				return
			}
		}
		conn.Write([]byte("Command received\n"))
		dataCallbackErrCh <- nil // Indicate no errors in callback
	}

	onErrorCallback := func(err error, conn net.Conn) {
		// Pass the error to the error channel
		errCh <- err
	}

	// Start server in a goroutine
	go func() {
		defer wg.Done()
		err := Start(config, onDataCallback, onErrorCallback)
		if err != nil {
			errCh <- err
		} else {
			errCh <- nil
		}
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

	// Check for errors from the server
	select {
	case serverErr := <-errCh:
		if serverErr != nil {
			t.Fatalf("Server error: %v", serverErr)
		}
	default:
		// No error
	}

	// Check for errors from the data callback
	select {
	case dataErr := <-dataCallbackErrCh:
		if dataErr != nil {
			t.Errorf("Data callback error: %v", dataErr)
		}
	default:
		// No error
	}
}

