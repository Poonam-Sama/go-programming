package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	//"strconv"
	"os"
)

func main() {
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		output, _ := bufio.NewReader(c).ReadString('\n')
		output = strings.TrimSuffix(output, "\n")
		res2 := strings.Split(output, ",")
		for i := 1; i < len(res2); i++ {
			fmt.Println(res2[i])
		}
		if res2[0] == "100" {
			return
		}
		var option string
		fmt.Scanln(&option)
		fmt.Fprintf(c, option+"\n")

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
