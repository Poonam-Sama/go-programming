package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Main function
type user struct {
	my_balance      int
	txn_times_limit int
}

func main() {
	user_1 := user{
		my_balance:      9563,
		txn_times_limit: 5,
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("Server has been listening..")
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	globalWithdraw := 0
	previousState := 0
	state := 0
	netData := bufio.NewReader(c)
	previous := ""
	for {
		switch state {
		case 0:
			{
				send := "0," + previous + "Enter 1 for checking balance,Enter 2 for withdrawl,Anyother key to exit"
				previous = ""
				previousState = 0
				c.Write([]byte(send + "\n"))
				option, _ := netData.ReadString('\n')
				option = strings.TrimSpace(string(option))
				if option == "1" {
					state = 1
				} else if option == "2" {
					state = 2
				} else {
					state = 100
				}
			}
		case 1:
			{
				previous = "Your balance is:" + strconv.Itoa(user_1.my_balance) + ","
				if previousState == 0 {
					state = 0
				} else {
					state = 11
				}
			}
		case 100:
			{
				send := "100," + "Exiting"
				c.Write([]byte(send + "\n"))
				c.Close()
				return
			}
		case 2:
			{
				send := "2," + previous + "Enter the amount in multiple of 100"
				c.Write([]byte(send + "\n"))
				amount, _ := netData.ReadString('\n')
				amount = strings.TrimSpace(string(amount))
				if _, err := strconv.Atoi(amount); err == nil {
					globalWithdraw, err = strconv.Atoi(amount)
				}
				if err != nil {
					previous = "Enter valid integer"
					state = 12
					continue
				}
				state = nextState(amount, user_1.my_balance)
			}
		case 4:
			{
				state = 12
				previous = "Not a valid Integer,"
			}
		case 5:
			{
				state = 12
				previous = "Max limit of withdraw is 5000,"
			}
		case 6:
			{
				state = 12
				previous = "Zero or negative amount is not permitted,"
			}
		case 7:
			{
				state = 12
				previous = "Amount is not multiple of 100,"
			}
		case 8:
			{
				state = 12
				previous = "Insufficient Balance,Your Balance is" + strconv.Itoa(user_1.my_balance) + ","
			}
		case 9:
			{ //When everything is correct
				state = 10
			}
		case 10:
			{
				user_1.my_balance = user_1.my_balance - globalWithdraw
				user_1.txn_times_limit = user_1.txn_times_limit - 1
				send := "0,Current Balance: " + strconv.Itoa(user_1.my_balance) + "," + print_note(globalWithdraw) + ","
				if user_1.my_balance < 100 || user_1.txn_times_limit == 0 {
					previous = send
					state = 11
					continue
				} else {
					for {
						send = send + ",Press Y to continue withdrawl or anyother key to exit"
						c.Write([]byte(send + "\n"))
						previous = ""
						cont, _ := netData.ReadString('\n')
						cont = strings.TrimSpace(string(cont))
						if cont == "Y" || cont == "y" {
							state = 2
							break
						} else {
							state = 0
							break
						}
					}
				}

			}
		case 11:
			{
				send := "0," + previous + "Enter 1 for checking balance,Anyother key to exit"
				previous = ""
				previousState = 11
				c.Write([]byte(send + "\n"))
				option, _ := netData.ReadString('\n')
				option = strings.TrimSpace(string(option))
				if option == "1" {
					state = 1
				} else {
					state = 100
				}
			}
		case 12:
			{
				send := "0," + previous + "Press Y to continue withdrawl or anyother key to exit"
				c.Write([]byte(send + "\n"))
				previous = ""
				cont, _ := netData.ReadString('\n')
				cont = strings.TrimSpace(string(cont))
				if cont == "Y" || cont == "y" {
					state = 2
				} else {
					state = 0
				}

			}
		}

	}
}

func nextState(amount_wid_string string, amount int) int {
	amount_wid, err := strconv.Atoi(amount_wid_string)
	if err != nil {
		return 4
	}
	if amount_wid > 5000 {
		return 5
	}

	if amount_wid <= 0 {
		return 6
	}
	if amount_wid%100 != 0 {
		return 7
	}
	if amount < amount_wid {
		return 8
	} else {
		return 9
	}

}

func valid_amount(amount_wid int) string {

	if amount_wid > 5000 {
		fmt.Println("Max limit of withdraw is 5000")
		return "1,1,Max limit of withdraw is 5000,Please Enter the amount in multiple of 100"
	}

	if amount_wid <= 0 {
		fmt.Println("Zero or negative amount is not permitted")
		return "1,1,Zero or negative amount is not permitted,Please Enter the amount in multiple of 100"
	}
	if amount_wid%100 != 0 {
		fmt.Println("Amount is not multiple of 100")
		return "1,1,Amount is not multiple of 100,Please Enter the amount in multiple of 100"
	} else {
		return "0"
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

	return "Note detais : " + xIn + "*500+" + yIn + "*200+" + zIn + "*100"

}
