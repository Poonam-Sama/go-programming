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

					fmt.Println("Enter Y for continue or any other key to exit")
					var option_2 string
					fmt.Scanln(&option_2)
					switch {
					case option_2 == "Y" || option_2 == "y":
						pass = 1

					default:
						pass = -1

					}

				} else {

					if validateAmount(user_1, val) {
						if user_1.txn_times_limit > 0 {

							printDenominations(val)
							user_1.my_balance = user_1.my_balance - val
							fmt.Println("Your remaining balance is ", user_1.my_balance)

							user_1.txn_times_limit--
							fmt.Println("Number of transaction left for a day: ", user_1.txn_times_limit)

							pass_1 := 1

							if user_1.txn_times_limit == 0 {
								pass = -1
								pass_1 = -1
							}
							for pass_1 > 0 && user_1.txn_times_limit > 0 {
								if user_1.my_balance > 100 {
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

									}

								}
								if user_1.my_balance < 100 {
									fmt.Println("Your balance is less than minimum balance")
									pass_1 = -1
									pass = -1

								}

							}

						} else {
							fmt.Println("Per day transaction limit exceeded")
							pass = -1
						}

					} else {
						fmt.Println("Enter Y for continue or any other key to exit")
						var option_1 string
						fmt.Scanln(&option_1)
						switch {
						case option_1 == "Y" || option_1 == "y":
							pass = 1

						default:
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
	fmt.Println("Thank You for choosing us.")

}

func printDenominations(amount_wid int) {

	var x, y, z int
	x = amount_wid / 500
	y = (amount_wid - (x * 500)) / 200
	z = (amount_wid - (x*500 + y*200)) / 100

	fmt.Printf("Your transaction is successful\nNote detais : %d*500+%d*200+%d*100\n", x, y, z)

}

func validateAmount(r user, amount_wid int) bool {

	if amount_wid > 5000 {
		fmt.Println("Maximum limit of withdraw is 5000.")
		return false
	}

	if amount_wid <= 0 {
		fmt.Println("Zero or negative amount is not permitted.")
		return false
	}

	if amount_wid%100 != 0 {
		fmt.Println("Amount is not multiple of 100.")
		return false
	}

	if amount_wid > r.my_balance {
		fmt.Println("This amount is unsufficient to withdrawal.\nPlease enter amount less than your current balance.")
		return false
	} else {
		return true
	}

}
