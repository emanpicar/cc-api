package main

import (
	"fmt"

	"github.com/emanpicar/cc-api/card"
	"github.com/emanpicar/cc-api/luhnalg"
)

func main() {
	luhnManager := luhnalg.New()
	cardManager := card.New(luhnManager)

	fmt.Printf("VISA %v \n", cardManager.GetCardDetails("4111111111111111"))
	fmt.Printf("VISA %v \n", cardManager.GetCardDetails("4111111111111"))
	fmt.Printf("VISA %v \n", cardManager.GetCardDetails("4012888888881881"))
	fmt.Printf("AMEX %v \n", cardManager.GetCardDetails("378282246310005"))
	fmt.Printf("Discover %v \n", cardManager.GetCardDetails("6011111111111117"))
	fmt.Printf("MasterCard %v \n", cardManager.GetCardDetails("5105105105105100"))
	fmt.Printf("MasterCard %v \n", cardManager.GetCardDetails("5105 1051 0510 5106"))
	fmt.Printf("Unknown %v \n", cardManager.GetCardDetails("9111111111111111"))

	fmt.Println("##########################")
	fmt.Printf("############# %v", cardManager.GetCardListDetails([]string{
		"4111111111111111",
		"4111111111111",
		"4012888888881881",
		"378282246310005",
		"6011111111111117",
		"5105105105105100",
		"5105 1051 0510 5106",
		"9111111111111111",
	}))
}
