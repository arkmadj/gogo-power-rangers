package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - Add item, s - Save bill, t - Add tip):", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price: ", reader)
		fmt.Println(name, price)
	case "s":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "t":
		fmt.Println("You chose t")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)
}
