package main

import (
	"bufio"
	"fmt"
	"os"
)

/* type person struct {
	Name         string
	Age          int
	Address      string
	Phone_Number string
	Country_Code int
}
*/
func username1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name : ")
	Name, _ := reader.ReadString('\n')

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address : ")
	Address, _ := reader1.ReadString('\n')

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")
	Age, _ := reader2.ReadString('\n')

	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your phone number")
	Phone_Number, _ := reader3.ReadString('\n')

	reader4 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the country code")
	Country_Code, _ := reader4.ReadString('\n')

	fmt.Println(Name)
	fmt.Println(Age)
	fmt.Println(Address)
	fmt.Println(Phone_Number)
	fmt.Println(Country_Code)

}
func main() {

	username1()

}
