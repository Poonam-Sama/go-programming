package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var response string

type amountJson struct {
	Input int `json:"amount" `
}

const (
	database    = "USER_DATABASE"
	collection1 = "BALANCE_COLLECTION"
	collection2 = "WITHDRAWAL_COLLECTION"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id"`
	Balance int           `bson:"balance"`
}

type withdrawalAmount struct {
	ID               bson.ObjectId `bson:"_id"`
	UserID           bson.ObjectId `bson:"userID"`
	WithdrawalAmount int           `bson:"amount"`
	TimeStamp        time.Time     `bson:"timestamp"`
}

func main() {
	Host := []string{
		"127.0.0.1:27017",
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	defer session.Close()
	balanceTable := session.DB(database).C(collection1)
	withdrawalTable := session.DB(database).C(collection2)

	r := gin.Default()
	r.GET("/balance/:id", func(c *gin.Context) {
		isClientIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if isClientIDvalid == false {
			c.JSON(400, gin.H{"Result": "Invalid Object ID"})
		} else {
			balanceID := bson.ObjectIdHex(c.Param("id"))
			fmt.Println(balanceID)
			var result userBalance
			err1 := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).One(&result)
			if err1 != nil {
				panic(err1)
			}
			c.JSON(200, gin.H{"Your balance is": result.Balance})

		}

	})

	r.POST("/withdrawal/:id", func(c *gin.Context) {
		isClientIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if isClientIDvalid == false {
			c.JSON(400, gin.H{"Result": "Invalid Object ID"})
		} else {

			var amountWithdrawal amountJson
			c.BindJSON(&amountWithdrawal)
			fmt.Println(amountWithdrawal.Input)
			balanceID := bson.ObjectIdHex(c.Param("id"))
			c.JSON(200, gin.H{
				"result": makeTransaction(amountWithdrawal.Input, balanceID, balanceTable, withdrawalTable)})
			response = ""

		}
	})

	r.Run()

}

func makeTransaction(amountWithdrawal int, balanceID bson.ObjectId, balanceTable *mgo.Collection, withdrawalTable *mgo.Collection) string {
	var result userBalance
	err1 := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).One(&result)
	if err1 != nil {
		panic(err1)
	}
	balance := result.Balance
	presentTime := time.Now()
	timeDifference := presentTime.Add(time.Hour * (-24))
	fmt.Println(presentTime, timeDifference)
	transactionCountLimit, err := withdrawalTable.Find(bson.M{"userID": balanceID, "timestamp": bson.M{"$gt": timeDifference}}).Count()
	fmt.Println("HI", transactionCountLimit)
	if err != nil {
		return "Internal error in transaction:" + err.Error()
	}
	if transactionCountLimit >= 5 {
		response += " " + "Per day transaction attempts limit exceeded"
		return response

	}

	errorCode, amountWithdrawal := amountValidate(amountWithdrawal, balance)

	switch errorCode {
	case 0:
		response += " " + "Maximum limit of withdrawal is 5000."

	case 1:
		response += " " + "Negative amount is not permitted."

	case 2:
		response += " " + "Amount is not multiple of 100."

	case 3:
		response += " " + "Transaction can not be made because given amount is more than account balance."
	case 10:

		response += " " + "Please enter a natural number"

	default:

		balance = balance - amountWithdrawal
		withdrawalID := bson.NewObjectId()
		if err := withdrawalTable.Insert(&withdrawalAmount{ID: withdrawalID, UserID: balanceID, WithdrawalAmount: amountWithdrawal, TimeStamp: time.Now()}); err != nil {
			fmt.Println("Error", err)
		}

		selector := bson.M{"_id": balanceID}
		updator := bson.M{"$inc": bson.M{"balance": -amountWithdrawal}}
		if err := balanceTable.Update(selector, updator); err != nil {
			fmt.Println("Error in updating balance:", err)
		}
		x := printDenominations(amountWithdrawal)
		response += " " + x
		y := fmt.Sprintf("Your remaining balance is %d.", balance)
		response += " " + y
		z := fmt.Sprintf("Remaining attempts for transactions is %d", (4 - transactionCountLimit))
		response += " " + z

	}
	return response

}
func printDenominations(amountWithdrawal int) string {
	var x, y, z int
	x = amountWithdrawal / 500
	y = (amountWithdrawal - (x * 500)) / 200
	z = (amountWithdrawal - (x*500 + y*200)) / 100
	i := fmt.Sprintf("Your transaction is successful.Note details : %d*500+%d*200+%d*100.", x, y, z)
	fmt.Println(i)
	return i
}

func amountValidate(amountWithdrawal int, balance int) (int, int) {
	if amountWithdrawal == 0 {
		fmt.Println("Please enter a natural number")
		return 10, amountWithdrawal
	}
	if amountWithdrawal > 5000 {
		fmt.Println("Maximum limit of withdrawal is 5000.")
		return 0, amountWithdrawal
	}

	if amountWithdrawal < 0 {
		fmt.Println("Negative amount is not permitted.")
		return 1, amountWithdrawal
	}

	if amountWithdrawal%100 != 0 {
		fmt.Println("Amount is not multiple of 100.")
		return 2, amountWithdrawal
	}

	if amountWithdrawal > balance {
		fmt.Println("Transaction can not be made because given amount is more than account balance. ")
		return 3, amountWithdrawal
	} else {
		return 4, amountWithdrawal
	}
}
