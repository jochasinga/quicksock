quicksock
=========

A quick example Go package to quickly create an instance of TCP socket client and server.

Install
-------
```Shell

$ go get https://github.com/jochasinga/quicksock.git

```

Examples
--------

Creating a TCP client:

```Go

package main

import qs "github.com/jochasinga/quicksock"

func main() {
        qs.TCPClient("google.com:www")
}

```

Creating a TCP server:

```Go

package main

import qs "github.com/jochasinga/quicksock"

func main() {
        qs.TCPServer("8080")

```