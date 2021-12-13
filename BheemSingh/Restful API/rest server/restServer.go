package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var balance = 9563
var txn_times_limit = 5

func get(w http.ResponseWriter, r *http.Request) {

	output := fmt.Sprintf("Your balance is %d", balance)
	w.Write([]byte(output))
}

func post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("ERROR")
	}

	var amount_wid int
	err = json.Unmarshal(body, &amount_wid)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Not a valid integer"))
		return

	}

	if txn_times_limit < 1 {
		w.Write([]byte("Per day transaction attempts limit exceeded"))
		return
	}

	error_code, amount_wid := amt_valid(amount_wid, balance)

	switch error_code {
	case 0:
		w.Write([]byte("Maximum limit of withdraw is 5000."))
		break
	case 1:
		w.Write([]byte("Zero or negative amount is not permitted."))
		break
	case 2:
		w.Write([]byte("Amount is not multiple of 100."))
		break
	case 3:
		w.Write([]byte("This amount is unsufficient to withdrawal."))
		break
	default:

		balance = balance - amount_wid
		x := printDenominations(amount_wid)
		w.Write([]byte(x))
		y := fmt.Sprintf("Your remaining balance is %d\n", balance)
		w.Write([]byte(y))
		txn_times_limit--
		z := fmt.Sprintf("Remaining attempts for transactions is %d", txn_times_limit)
		w.Write([]byte(z))
		break

	}

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

func amt_valid(amount_wid int, balance int) (int, int) {
	if amount_wid > 5000 {
		fmt.Println("Maximum limit of withdraw is 5000.")
		return 0, amount_wid
	}

	if amount_wid < 0 || amount_wid == 0 {
		fmt.Println("Zero or negative amount is not permitted.")
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
func put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Program terminated"))
	os.Exit(1)
}

func requests(w http.ResponseWriter, r *http.Request) {
	message := "[{'request':GET,'end-point':'/balance','usage':'For getting balance'},\n{'request':'POST','end-point':'/withdraw','usage':'for withdrawing money'},\n{'request':'PUT','end-point':'/exit','usage':'For terminating program'}]"
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	switch r.Method {
	case "GET":
		if r.URL.Path == "/balance" {
			get(w, r)
			w.Write([]byte("\n" + message))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		w.Write([]byte("\n" + message))
	case "POST":
		if r.URL.Path == "/withdraw" {
			post(w, r)
			w.Write([]byte("\n" + message))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		w.Write([]byte("\n" + message))
	case "PUT":
		if r.URL.Path == "/exit" {
			put(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		w.Write([]byte("\n" + message))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		w.Write([]byte("\n" + message))
		return
	}
}

func main() {
	http.HandleFunc("/", requests)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
