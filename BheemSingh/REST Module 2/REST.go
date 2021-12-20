package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var balance = 9563
var txn_times_limit = 5
var getbody = ""
var exit = 0

func get(API [3]string, conn net.Conn) {

	output := fmt.Sprintf("Your balance is %d", balance)
	// conn.Write([]byte(output))
	getbody += " " + output
}

func post(API [3]string, conn net.Conn) {

	var amount_wid int

	amount_wid, err := strconv.Atoi(API[2])
	if err != nil {
		fmt.Println(err)

		getbody += " " + "Not a valid integer"
		return

	}

	if txn_times_limit < 1 {
		getbody += " " + "Per day transaction attempts limit exceeded"
		return
	}

	error_code, amount_wid := amt_valid(amount_wid, balance)

	switch error_code {
	case 0:
		getbody += " " + "Maximum limit of withdraw is 5000."
		break
	case 1:
		getbody += " " + "Zero or negative amount is not permitted."
		break
	case 2:
		getbody += " " + "Amount is not multiple of 100."
		break
	case 3:
		getbody += " " + "This amount is unsufficient to withdrawal."
		break
	default:

		balance = balance - amount_wid
		x := printDenominations(amount_wid)
		getbody += " " + x
		y := fmt.Sprintf("Your remaining balance is %d\n", balance)
		getbody += " " + y
		txn_times_limit--
		z := fmt.Sprintf("Remaining attempts for transactions is %d", txn_times_limit)
		getbody += " " + z
		break

	}

}
func printDenominations(amount_wid int) string {
	var x, y, z int
	x = amount_wid / 500
	y = (amount_wid - (x * 500)) / 200
	z = (amount_wid - (x*500 + y*200)) / 100
	i := fmt.Sprintf("Your transaction is successful\nNote detais : %d*500+%d*200+%d*100\n", x, y, z)
	fmt.Println(i)
	return i
}

func amt_valid(amount_wid int, balance int) (int, int) {
	if amount_wid > 5000 {
		fmt.Println("Maximum limit of withdraw is 5000.")
		return 0, amount_wid
	}

	if amount_wid < 0 || amount_wid == 0 {
		fmt.Println("Zero or negative amount is not permitted.")
		return 1, amount_wid
	}

	if amount_wid%100 != 0 {
		fmt.Println("Amount is not multiple of 100.")
		return 2, amount_wid
	}

	if amount_wid > balance {
		fmt.Println("Amount withdrawl is less than current balance.")
		return 3, amount_wid
	} else {
		return 4, amount_wid
	}
}

func put(conn net.Conn) {
	getbody += " " + "Program has been terminated"
	respond(conn)
	os.Exit(1)
}

func requests(API [3]string, conn net.Conn) {
	message := "[{'request':GET,'end-point':'/balance','usage':'For getting balance'},\n{'request':'POST','end-point':'/withdraw','usage':'for withdrawing money'},\n{'request':'PUT','end-point':'/exit','usage':'For terminating program'}]"
	fmt.Println(API[0])
	fmt.Println(API[1])
	switch API[0] {
	case "GET":
		if API[1] == "/balance" {
			get(API, conn)
			return
		}
		getbody += " " + "method not allowed" + "\n" + message

	case "POST":
		if API[1] == "/withdraw" {
			post(API, conn)

			return
		}
		getbody += " " + "method not allowed" + "\n" + message

	case "PUT":
		if API[1] == "/exit" {
			put(conn)
			return
		}

		getbody += " " + "method not allowed" + "\n" + message

	default:

		getbody += " " + "method not allowed" + "\n" + message

		return
	}
}

func respond(conn net.Conn) {
	body := getbody
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func API_Detail(conn net.Conn) [3]string {

	scanner := bufio.NewScanner(conn)
	//fmt.Println(scanner)
	i := 0
	t := 0
	API_Object := [3]string{}
	for scanner.Scan() {
		input := scanner.Text()
		if i > 1 {
			if API_Object[0] == "PUT" || API_Object[0] == "GET" {
				return API_Object
			}
		}
		fmt.Println(input)
		if i == 0 {
			API_Object[0] = strings.Fields(input)[0]
			API_Object[1] = strings.Fields(input)[1]
		}
		if t == 1 {
			body := input
			API_Object[2] = body
			return API_Object
		}
		if input == "" {
			t = 1
		}
		fmt.Println(i)
		i++

	}
	return API_Object
}

func main() {
	fmt.Println("start main")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		getbody = ""
		API := API_Detail(conn)
		requests(API, conn)
		respond(conn)
		conn.Close()

	}

}
