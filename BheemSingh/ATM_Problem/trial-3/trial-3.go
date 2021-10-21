package main

import (
	"fmt"
	"strconv"
)

type user struct {
	my_balance int
	// txn_times_limit means per day limit of transaction times
	txn_times_limit int
}

func main() {

	user_1 := user{
		my_balance:      9563,
		txn_times_limit: 5,
	}
	var option string

	// amount_wid means amount to be withdrawal
	var amount_wid string

	for i := 1; i > 0; i++ {

		x := "Press 1 : for balance check"
		fmt.Println(x)
		if user_1.txn_times_limit > 0 && user_1.my_balance > 100 {
			y := "Press 2 : Withdraw money"
			fmt.Println(y)
		}
		z := "Press any other key : for exit "
		fmt.Println(z)
		fmt.Scanln(&option)
		switch {
		case option == "1":
			fmt.Println("Your current balance is", user_1.my_balance)
			option = "0"

		case (option == "2" && user_1.txn_times_limit > 0 && user_1.my_balance > 100):
			pass := 1
			for pass > 0 {

				fmt.Println("Enter the amount you want to withdraw ")

				fmt.Scanln(&amount_wid)
				val, err := strconv.Atoi(amount_wid)

				if err != nil {
					fmt.Printf("%s is not a valid integer\n", amount_wid)
				} else {
					// fmt.Printf("%s correct input %d\n", amount_wid, val)

					if valid_amount(&user_1, val) {
						if user_1.txn_times_limit > 0 {
							withdraw(&user_1, val)
							user_1.txn_times_limit--
							fmt.Println("Number of trsnsaction left for a day: ", user_1.txn_times_limit)

							pass_1 := 1
							count := 0
							if user_1.txn_times_limit == 0 {
								pass = -1
								pass_1 = -1
							}
							for pass_1 > 0 && user_1.txn_times_limit > 0 {

								var choose string
								fmt.Println("Do you want to withdraw again?")
								fmt.Println("choose Y for yes and N for no")
								fmt.Scanln(&choose)

								switch {
								case choose == "Y" || choose == "y":
									pass = 1
									pass_1 = -1

								case choose == "N" || choose == "n":
									pass = -1
									pass_1 = -1

								default:
									fmt.Println("Please enter valid keyword")
									count++
									// we are allowing maximum three wrong inputs.
									if count == 3 {
										pass_1 = -1
										pass = -1
									}

								}

							}

						} else {
							fmt.Println("Per day transaction limit exceeded")
							pass = -1
						}

					}
				}

			}

			option = "0"
		default:
			i = -1
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

	fmt.Println("Your remaining balance is ", x.my_balance)

}

func valid_amount(r *user, amount_wid int) bool {
	min_balance := 100

	if amount_wid > 5000 {
		fmt.Println("Max limit of withdraw is 5000")
		return false
	}
	if amount_wid > (r.my_balance - min_balance) {
		if r.my_balance >= amount_wid {
			fmt.Println("Withdrawl can't be made as minimum balance in your account should be 100 rupees")
		} else {
			fmt.Println("Balance is insufficient")
		}
		return false
	}
	if amount_wid < 0 {
		fmt.Println("Negative amount is not permitted")
		return false
	}
	if amount_wid%100 != 0 {
		fmt.Println("Amount is not multiple of 100")
		return false
	} else {
		return true
	}
}
