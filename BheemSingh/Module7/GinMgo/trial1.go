package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var getbody string

type amountJson struct {
	Input int `json:"amount" `
}

const (
	database    = "DATABASE_1"
	collection1 = "BAL_COLL_1"
	collection2 = "WID_COLL_1"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id"`
	Balance int           `bson:"balance"`
	//txn_times_limit int           `bson:"count"`
}

type withdrwalAmount struct {
	ID         bson.ObjectId `bson:"_id"`
	UserID     bson.ObjectId `bson:"ID"`
	Wid_Amount int           `bson:"amount"`
	TimeStamp  time.Time     `bson:"Timestamp"`
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

			var amount_wid amountJson
			c.BindJSON(&amount_wid)
			fmt.Println(amount_wid.Input)
			balanceID := bson.ObjectIdHex(c.Param("id"))
			c.JSON(200, gin.H{
				"result": makeTransaction(amount_wid.Input, balanceID, balanceTable, withdrawalTable)})
			getbody = ""

		}
	})

	r.Run()

}

func makeTransaction(amount_wid int, balanceID bson.ObjectId, balanceTable *mgo.Collection, withdrawalTable *mgo.Collection) string {
	var result userBalance
	err1 := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).One(&result)
	if err1 != nil {
		panic(err1)
	}
	balance := result.Balance
	presentTime := time.Now()
	timeDifference := presentTime.Add(time.Hour * (-24))
	fmt.Println(presentTime, timeDifference)
	txn_times_limit, err := withdrawalTable.Find(bson.M{"ID": balanceID, "Timestamp": bson.M{"$gt": timeDifference}}).Count()
	if err != nil {
		return "Internal error in transaction:" + err.Error()
	}
	if txn_times_limit >= 5 {
		getbody += " " + "Per day transaction attempts limit exceeded"
		return getbody

	}

	errorCode, amount_wid := amountValidate(amount_wid, balance)

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

		balance = balance - amount_wid
		wid_id := bson.NewObjectId()
		if err := withdrawalTable.Insert(&withdrwalAmount{ID: wid_id, UserID: balanceID, Wid_Amount: amount_wid, TimeStamp: time.Now()}); err != nil {
			fmt.Println("Error", err)
		}

		selector := bson.M{"_id": balanceID}
		updator := bson.M{"$inc": bson.M{"balance": -amount_wid}}
		if err := balanceTable.Update(selector, updator); err != nil {
			fmt.Println("Error in updating balance:", err)
		}
		x := printDenominations(amount_wid)
		getbody += " " + x
		y := fmt.Sprintf("Your remaining balance is %d\n", balance)
		getbody += " " + y
		txn_times_limit--
		// z := fmt.Sprintf("Remaining attempts for transactions is %d", txn_times_limit)
		// getbody += " " + z

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
