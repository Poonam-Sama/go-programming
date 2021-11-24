package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
		if message == "exit" {

		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose the option ")
		option, _ := reader.ReadString('\n')
		fmt.Fprint(conn, option+"\n")
	}
	//result, err := ioutil.ReadAll(conn)
	//fmt.Println(string(result))
	checkError(err)
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
