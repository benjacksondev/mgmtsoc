# MgmtSoc

[![Go](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml/badge.svg)](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml)


MgmtSoc is a simple TCP server library in Go that allows you to start a server with configurable host and port. It handles incoming data and errors through user-defined callback functions.

## Description

Inspired by this [blog post](https://www.marccostello.com/little-socket-services/), MgmtSoc provides a lightweight alternative to building HTTP APIs for interacting with long-running processes, such as daemons or other background services. Traditionally, setting up such interactions requires a setting up a 

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
		fmt.Printf("Command: %s, Args: %v\n", cmd, args)
		conn.Write([]byte("Command received\n"))
	}

	onErrorCallback := func(err error, conn net.Conn) {
		fmt.Println("Error:", err)
	}

	success := mgmtsoc.Start(config, onDataCallback, onErrorCallback)
	if success {
		fmt.Println("Server started successfully")
	}
}
```

## Testing
To run the tests for this library, use the `go test` command:

```bash
go test
```

