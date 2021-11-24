package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type user struct {
	my_balance      int
	txn_times_limit int
}

func main() {

	user_1 := user{
		my_balance:      9563,
		txn_times_limit: 5,
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// service := ":7777"
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	// checkError(err)

	// listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("continue")
			continue
		}

		for {
			netData, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return

			}
			option := strings.TrimSpace(string(netData))
			fmt.Println("Input from client : ", string(netData))
			// fmt.Println(y)

			switch {
			case option == "balance":
				balance := strconv.FormatInt(int64(user_1.my_balance), 10)
				//	fmt.Println(balance)
				fmt.Fprint(conn, balance+"\n")
			case option == "count":
				count := strconv.FormatInt(int64(user_1.txn_times_limit), 10)
				fmt.Fprint(conn, count+"\n")

			case option == "1":
				balance := strconv.FormatInt(int64(user_1.my_balance), 10)
				fmt.Println(balance)
				fmt.Fprint(conn, balance+"\n")
			case option == "2":
				amt_withdrawl, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				amt_withdrawl = strings.TrimSpace(string(amt_withdrawl))
				amt_int, err := strconv.Atoi(amt_withdrawl)
				var status string
				if err != nil {
					fmt.Printf("%s is not a valid integer\n", amt_withdrawl)
					status = "400"
				} else {
					status = valid_amount(amt_int)
				}
				fmt.Fprint(conn, status+"\n")
				if status == "0" {
					user_1.my_balance = user_1.my_balance - amt_int
					fmt.Println("Your remaining balance is ", user_1.my_balance)
					user_1.txn_times_limit--
					fmt.Println("Number of trsnsaction left for a day: ", user_1.txn_times_limit)
					fmt.Fprint(conn, strconv.FormatInt(int64(user_1.my_balance), 10)+"\n")
					denom := print_note(amt_int)
					fmt.Fprint(conn, denom+"\n")

				}

			default:
				conn.Write([]byte("exit"))
				conn.Close()
				return

			}
		}
	}

}

func print_note(amount_wid int) string {

	var x, y, z int
	x = amount_wid / 500
	y = (amount_wid - (x * 500)) / 200
	z = (amount_wid - (x*500 + y*200)) / 100
	xIn := strconv.FormatInt(int64(x), 10)
	yIn := strconv.FormatInt(int64(y), 10)
	zIn := strconv.FormatInt(int64(z), 10)

	return "Note detais : " + xIn + "*500+" + yIn + "*200+" + zIn + "*100\n"

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func valid_amount(amount_wid int) string {

	if amount_wid > 5000 {
		fmt.Println("Max limit of withdraw is 5000")
		return "100"
	}

	if amount_wid <= 0 {
		fmt.Println("Zero or negative amount is not permitted")
		return "200"
	}

	if amount_wid%100 != 0 {
		fmt.Println("Amount is not multiple of 100")
		return "300"
	} else {
		return "0"
	}

}
