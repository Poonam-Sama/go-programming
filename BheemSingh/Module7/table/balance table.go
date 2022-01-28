package main

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	type userBalance struct {
		ID      bson.ObjectId `bson:"_id"`
		Balance int           `bson:"balance"`
	}

	Host := []string{
		"127.0.0.1:27017",
	}
	const (
		Username    = "YOUR_USERNAME"
		Password    = "YOUR_PASS"
		myDatabase  = "DATABASE_1"
		Collection1 = "BAL_COLL_1"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("error", err)
	}

	coll := session.DB(myDatabase).C(Collection1)
	fmt.Println("Collection type:", reflect.TypeOf(coll), "\n")
	fmt.Println(coll.RemoveAll(nil))
	id1 := bson.NewObjectId()
	if err := coll.Insert(&userBalance{ID: id1, Balance: 9563}); err != nil {
		fmt.Println("error", err)
	}
	var result []userBalance

	err3 := coll.Find(nil).All(&result)
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance: ", result)
	}

}
