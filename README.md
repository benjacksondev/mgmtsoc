# MgmtSoc

[![Go](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml/badge.svg)](https://github.com/benjacksondev/mgmtsoc/actions/workflows/ci.yml)


MgmtSoc is a simple TCP server library in Go that allows you to start a server with configurable host and port. It handles incoming data and errors through user-defined callback functions.

## Description

Inspired by Marc Costello's blog post, MgmtSoc provides a lightweight alternative to building complex HTTP APIs for interacting with long-running processes, such as daemons or other background services. Traditionally, setting up such interactions requires an HTTP framework, which involves additional overhead and documentation. MgmtSoc simplifies this process by leveraging plain TCP socs, offering a straightforward and efficient solution.

Marc Costello shared his insights on using TCP socs instead of HTTP APIs for management purposes. He noted how Etsy's `statsd` project employed this technique, which inspired the creation of MgmtSoc. This library allows developers to implement management socs in any language, providing a simple and effective method for communication with long-running processes.

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

