package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database    = "DATABASE"
	collection1 = "BAL_COLL"
	collection2 = "WID_COLL"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id"`
	Balance int           `bson:"balance" `
}

type withdrwalAmount struct {
	ID         bson.ObjectId `bson:"_id"`
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

	var amount_wid string
	fmt.Println("Enter the amount you want to withdraw ")
	fmt.Scanln(&amount_wid)
	val, err := strconv.Atoi(amount_wid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	balanceTable := session.DB(database).C(collection1)
	withdrawalTable := session.DB(database).C(collection2)
	balanceTable.RemoveAll(nil)
	//withdrawalTable.RemoveAll(nil)

	balance_id := bson.NewObjectId()
	if err := balanceTable.Insert(&userBalance{ID: balance_id, Balance: 9563}); err != nil {
		fmt.Println("Error", err)
	}

	var result userBalance
	err1 := balanceTable.Find(bson.M{"_id": balance_id}).One(&result)
	if err1 != nil {
		panic(err1)
	}
	var x int
	x = result.Balance
	fmt.Println("balance is ", x)
	if validateAmount(val, x) {
		wid_id := bson.NewObjectId()
		if err := withdrawalTable.Insert(&withdrwalAmount{ID: wid_id, Wid_Amount: val, TimeStamp: time.Now()}); err != nil {
			fmt.Println("Error", err)
		}

		selector := bson.M{"_id": balance_id}
		updator := bson.M{"$inc": bson.M{"balance": -val}}
		if err := balanceTable.Update(selector, updator); err != nil {
			fmt.Println("Error in updating balance:", err)
		}
	}

	var wid_table []withdrwalAmount
	err2 := withdrawalTable.Find(nil).All(&wid_table)
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Withdrawal Table", wid_table)
	}

	var balance_table []userBalance
	err3 := balanceTable.Find(nil).All(&balance_table)
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance Table", balance_table)
	}

}
func validateAmount(amount_wid int, x int) bool {

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

	if amount_wid > x {
		fmt.Println("This amount is unsufficient to withdrawal.\nPlease enter amount less than your current balance.")
		return false
	} else {
		return true
	}

}
