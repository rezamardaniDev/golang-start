package main

import (
	"fmt"
)

// this struct for view other cart
type creditCard struct {
	cardNumber string
	expireDate string
	cvv2       string
	bankName   string
}

func main() {
	message := sayNo("reza")
	fmt.Println(message)

	cards := []creditCard{
		{cardNumber: "6037991458632145", expireDate: "01/04", cvv2: "478", bankName: "bluBank"},
		{cardNumber: "5892101234567890", expireDate: "02/05", cvv2: "123", bankName: "meliBank"},
		{cardNumber: "6273531122334455", expireDate: "03/06", cvv2: "987", bankName: "parsianBank"},
		{cardNumber: "6037709876543210", expireDate: "04/07", cvv2: "555", bankName: "samanBank"},
		{cardNumber: "6037991458632145", expireDate: "01/04", cvv2: "478", bankName: "samanBank"},
	}

	for _, item := range cards {
		fmt.Println(item.cardNumber)
	}

	lst1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 14, 15, 16, 17, 18, 19, 20}

	for index, item := range lst1 {
		fmt.Println("index:", index, "item:", item)
	}

	i := 0
	for i < 20 {
		fmt.Println("this is:", i)
		i++
	}

	userDisplay("reza")
	sayHello("reza")

	isValid := cardChecker(cards[0])
	fmt.Println("Is the card valid?", isValid)
}

func sayHello(name string) {
	fmt.Printf("Hello %s", name)
}

func userDisplay(name string) {
	fmt.Printf("name: %s", name)
}

func cardChecker(card creditCard) bool {
	if len(card.cardNumber) != 16 || len(card.expireDate) != 5 || len(card.cvv2) != 3 {
		return false
	} else {
		return true
	}
}
