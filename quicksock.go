package quicksock

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func TCPClient(url string) (*net.TCPConn, error) {
	fmt.Printf("Connecting to %s", url)

	addr, err := net.ResolveTCPAddr("tcp4", url)
	checkError(err)
	
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)
	
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	return conn, nil
}

func TCPServer(port string) (*net.TCPListener, error) {

	addr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	
	listener, err := net.ListenTCP("tcp", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		greeting := "Hey this is server!"
		conn.Write([]byte(greeting))
		defer conn.Close()
	}

	return listener, nil
	
}

func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
