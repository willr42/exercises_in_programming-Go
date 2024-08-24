package exercise10

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

type Item struct {
	name  string
	price int
	qty   int
}

var taxRate = 5.5 / 100

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the shopping list app.")
	fmt.Println("Input an item, its price and its quantity.")
	fmt.Println("Input nothing to calculate your total.")

	var list []Item
	var total = 0

	for {
		name, err := util.GetInput(r, "Name of item: ")

		if name == "" {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		price, err := util.GetCentsWithPrompt(r, "Price of item: ")

		if price == 0 {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		qty, err := util.GetIntWithPrompt(r, "Quantity of item: ")

		if qty == 0 {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		total += price * qty
		list = append(list, Item{name, price, qty})

	}

	fmt.Printf("List: %v", list)
	totalStr := util.FormatCents(total)

	fmt.Printf("Subtotal: $%s\n", totalStr)

	totalTax := int(float64(total) * taxRate)

	taxStr := util.FormatCents(totalTax)
	fmt.Printf("Tax: $%s\n", taxStr)

	final := total + totalTax
	finalStr := util.FormatCents(final)
	fmt.Printf("Final total: $%s\n", finalStr)
}
