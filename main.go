package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Zultan's bill")

	myBill.format()
	fmt.Println(myBill.format())

}
