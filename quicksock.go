package quicksock

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func Client(protocol, host string, port int, msg string) (net.Conn, error) {

	fmt.Printf("Connecting to %s:%d", host, port)

	if protocol == "udp" {
		UDPClient(host, port, msg)
	}

	addr := host + ":" + string(port)

	conn, err := net.Dial(protocol, addr)
	checkError(err)

	_, err := conn.Write([]byte(msg))
	checkError(err)

	reply, err := iountil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(reply))

	return conn, nil
}

func UDPClient(host string, port int, msg string) (*net.UDPConn, error) {

	fmt.Printf("Sending packets to %s:%d", host, port)
	addr, err := net.ResolveUDPAddr("udp", (host + ":" + int(port)))
	checkError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)

	_, err = conn.Write([]byte(msg))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	return conn, nil

}

// Create a TCP Client that sends a HTTP request to a specified host
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

// Create a TCP server on the local host that listens on a specified port
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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
