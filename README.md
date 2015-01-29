quicksock
=========

A quick example Go package to quickly create an instance of TCP socket client and server.

Install
-------
```Shell

$ go get github.com/jochasinga/quicksock
$ go install github.com/jochasinga/quicksock

```

Examples
--------

Creating a general client request to google.com:

```Go

package main

import qs "github.com/jochasinga/quicksock"

func main() {
        qs.Client("tcp", "google.com", "www", "HEAD / HTTP/1.0\r\n\r\n")
}

```

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