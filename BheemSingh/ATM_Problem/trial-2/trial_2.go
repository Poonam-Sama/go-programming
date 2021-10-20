package main

import (
	"fmt"
	"strconv"
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
	var option string

	// amount_wid means amount to be withdraw
	var amount_wid string

	for i := 1; i > 0; i++ {

		x := "Press 1 : for balance check\nPress 2 : withdraw money\nPress any other key : for exit "
		fmt.Println(x)
		fmt.Scanln(&option)
		switch {
		case option == "1":
			fmt.Println("your current balance is", user_1.my_balance)

		case option == "2":
			fmt.Println("Enter the amount you want to withdraw ")

			fmt.Scanln(&amount_wid)
			val, err := strconv.Atoi(amount_wid)

			if err != nil {
				fmt.Printf("%s is not a valid input.\n", amount_wid)
			} else {
				// fmt.Printf("%s correct input %d\n", amount_wid, val)

				if valid_amount(&user_1, val) {
					if user_1.txn_times_limit > 0 {
						withdraw(&user_1, val)
						user_1.txn_times_limit--
						fmt.Println("Number of trsnsaction left for a day: ", user_1.txn_times_limit)
					} else {
						fmt.Println("per day transaction limit exceeded")
					}

				}
			}

		default:
			i = -9
		}
	}

}
func print_note(amount_wid int) {

	var x, y, z int
	x = amount_wid / 500
	y = (amount_wid - (x * 500)) / 200
	z = (amount_wid - (x*500 + y*200)) / 100

	fmt.Printf("Your transaction is successful\nNote detais : %d*500+%d*200+%d*100\n", x, y, z)
	fmt.Println("Thank you")
}

func withdraw(x *user, amount_wid int) {

	print_note(amount_wid)
	x.my_balance = x.my_balance - amount_wid

	fmt.Println("your remaining balance is ", x.my_balance)

}

func valid_amount(r *user, amount_wid int) bool {
	var min_balance int
	min_balance = 100

	if amount_wid > 5000 {
		fmt.Println("max limit of withdraw is 5000")
		return false
	}
	if amount_wid > (r.my_balance - min_balance) {
		fmt.Println("balance is insufficient")
		return false
	}
	if amount_wid < 0 {
		fmt.Println("Negative amount is not permitted")
		return false
	}
	if amount_wid%100 != 0 {
		if amount_wid < 100 {
			fmt.Println("min 100 rupees can be transacted")
		} else {
			fmt.Println("amount is not multiple of 100")
		}
		return false
	} else {
		return true
	}
}
