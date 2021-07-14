package main

import "fmt"

func main() {
	//fmt.Println("Hello" )
	minAmount := 5000
	purchase := 15000
	percentCashback := 15
	cashbackLimit := 2000
	cashback := 0
	const fullPercent = 100

	if purchase >= minAmount{
		cashback = purchase * percentCashback / fullPercent
		fmt.Println("Prediction cashback is", cashback)
	}
	if cashback > cashbackLimit{
		cashback = cashbackLimit
	}
	fmt.Println("cashback is", cashback)
}

//DRAFT
//git init
//git add .
//git commit -m "my message"
//git status