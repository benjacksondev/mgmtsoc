# MgmtSoc

[![Go](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml/badge.svg)](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml)


Lightweight TCP server library which allows you to start a server with configurable host and port. Handles incoming data and errors through user-defined callback functions.

tcp server 

## Installation

To install the package, run:


```bash
go get github.com/benjacksondev/mgmtsoc
```

## Usage

Example of how to use mgmtsoc:

```go
package main

import (
	"fmt"
	"net"

	"github.com/benjacksondev/mgmtsoc"
)

func main() {
	config := mgmtsoc.Config{MgmtHost: "localhost", MgmtPort: 8123}

	onDataCallback := func(cmd string, args []string, conn net.Conn) {
		if cmd == "status" {
			conn.Write([]byte("UP\n"))
		}

		if cmd == "help" {
			conn.Write([]byte("Commands: help status\n"))
		}
	}

	onErrorCallback := func(err error, conn net.Conn) {
		fmt.Println("Error:", err)
	}

	success := mgmtsoc.Start(config, onDataCallback, onErrorCallback)
	if success {
		fmt.Println("Server started successfully")
	}

	// Prevent the main function from exiting
	select {} // This blocks forever
}
```

### netcat (nc)

```
$ echo "help" | nc localhost 8123
Commands: help status
```

### telnet

```
$ telnet localhost 8123
> help
Commands: help status
```

## Testing
To run the tests for this library, use the `go test` command:

```bash
go test
```

