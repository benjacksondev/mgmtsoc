# MgmtSoc

[![Go](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml/badge.svg)](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml)


MgmtSoc is a simple TCP server library in Go that allows you to start a server with configurable host and port. It handles incoming data and errors through user-defined callback functions.

## Description

Needed a reuseable method to send commands to running processes, and inspired by this [blog post](https://www.marccostello.com/little-socket-services/), created a small go library which can be dropped into Go programs to do just that. Blog post refers to statsd by Etsy but, also seen this kind of thing done in Kafka services, zookeeper and alike.

## Installation

To install the package, run:


```bash
go get github.com/benjacksondev/mgmtsoc
```

## Usage

Here is an example of how to use the MgmtSoc library:

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

