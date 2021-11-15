package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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
		x := "Press 1 : for balance check"
		fmt.Println(x)
		// reader := bufio.NewReader(os.Stdin)
		optionBalance := "balance"
		fmt.Fprint(conn, optionBalance+"\n")
		balance, _ := bufio.NewReader(conn).ReadString('\n')
		balance = strings.TrimSuffix(balance, "\n")
		balanceInt, err := strconv.Atoi(balance)
		checkError(err)
		optionCount := "count"
		fmt.Fprint(conn, optionCount+"\n")
		count, _ := bufio.NewReader(conn).ReadString('\n')
		count = strings.TrimSuffix(count, "\n")
		countInt, err := strconv.Atoi(count)
		checkError(err)
		if countInt > 0 && balanceInt > 100 {
			y := "Press 2 : Withdraw money"
			fmt.Println(y)
		}

		z := "Press any other key : for exit "
		fmt.Println(z)

		var option string
		println(option)
		fmt.Scanln(&option)
		switch {
		case option == "1":
			fmt.Println("Your Balance is ", balance)
		case option == "2" && countInt > 0:
			for {
				fmt.Fprint(conn, "2"+"\n")
				fmt.Println("Enter the amount you want to withdrawl")
				var amt_withdrawl string
				fmt.Scanln(&amt_withdrawl)
				fmt.Fprint(conn, amt_withdrawl+"\n")

				status, _ := bufio.NewReader(conn).ReadString('\n')
				if status == "0\n" {
					current_balance, _ := bufio.NewReader(conn).ReadString('\n')
					denom, _ := bufio.NewReader(conn).ReadString('\n')
					fmt.Println("Your current Balance is: ", current_balance)
					fmt.Println(denom)
					var choose string
					optionCount := "count"
					fmt.Fprint(conn, optionCount+"\n")
					count, _ := bufio.NewReader(conn).ReadString('\n')
					fmt.Println("Remaining Withdrawl of the day: ", count)
					count = strings.TrimSuffix(count, "\n")
					countInt, err := strconv.Atoi(count)
					checkError(err)
					i := countInt > 0
					for i {
						fmt.Println("Do you want to withdraw again?")

						fmt.Println("choose Y for yes and N for no")

						fmt.Scanln(&choose)

						if choose == "N" || choose == "n" {
							i = false
						} else if choose == "Y" || choose == "y" {
							i = false
						} else {

							fmt.Println("Invalid Input!! \nPlease Choose from Y or N")
						}
					}
					if choose == "N" || choose == "n" || countInt <= 0 {
						break
					}
				}

				if status == "100\n" {
					fmt.Println("Max limit of withdraw is 5000")
				}
				if status == "200\n" {
					fmt.Println("Zero or negative amount is not permitted")
				}
				if status == "300\n" {
					fmt.Println("Amount is not multiple of 100")
				}
				if status == "400\n" {
					fmt.Println("Please enter valid amount")
				}

				if status == "100\n" || status == "200\n" || status == "300\n" || status == "400\n" {
					fmt.Println("Press Y for continue or any other key to exit")
					var option_1 string
					fmt.Scanln(&option_1)
					if option_1 == "Y" || option_1 == "y" {
						continue
					} else {
						break
					}

				}
			}

			continue
		default:
			conn.Close()
			return
		}

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
