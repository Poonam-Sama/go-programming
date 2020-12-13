package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*f,err:=os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()*/
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f) //to read from file f,which is names.txt(we hv to open it frst)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(bs)) //prints what it got frm the files
}
