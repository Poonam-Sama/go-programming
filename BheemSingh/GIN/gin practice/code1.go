package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var getbody string

type amountJson struct {
	Input int `json:"amount" `
}

func main() {
	var balance = 9563
	var txn_times_limit = 5

	r := gin.Default()
	r.GET("/balance", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Your current balance is ": balance,
		})
		getbody = ""
	})

	r.POST("/withdraw", func(c *gin.Context) {
		var amount_wid amountJson
		c.BindJSON(&amount_wid)
		fmt.Println(amount_wid.Input)

		c.JSON(200, gin.H{
			"result": makeTransaction(amount_wid.Input, &balance, &txn_times_limit)})
		getbody = ""
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func makeTransaction(amount_wid int, balance *int, txn_times_limit *int) string {

	if *txn_times_limit < 1 {
		getbody += " " + "Per day transaction attempts limit exceeded"
		return getbody

	}

	errorCode, amount_wid := amountValidate(amount_wid, *balance)

	switch errorCode {
	case 0:
		getbody += " " + "Maximum limit of withdraw is 5000."

	case 1:
		getbody += " " + "Negative amount is not permitted."

	case 2:
		getbody += " " + "Amount is not multiple of 100."

	case 3:
		getbody += " " + "This amount is unsufficient to withdrawal."
	case 10:

		getbody += " " + "Please enter a natural number"

	default:

		*balance = *balance - amount_wid
		x := printDenominations(amount_wid)
		getbody += " " + x
		y := fmt.Sprintf("Your remaining balance is %d\n", *balance)
		getbody += " " + y
		*txn_times_limit--
		z := fmt.Sprintf("Remaining attempts for transactions is %d", *txn_times_limit)
		getbody += " " + z

	}
	return getbody

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

func amountValidate(amount_wid int, balance int) (int, int) {
	if amount_wid == 0 {
		fmt.Println("Please enter a natural number")
		return 10, amount_wid
	}
	if amount_wid > 5000 {
		fmt.Println("Maximum limit of withdraw is 5000.")
		return 0, amount_wid
	}

	if amount_wid < 0 {
		fmt.Println("Negative amount is not permitted.")
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
