package balance

import (
	"fmt"
)

func calculateBalance(startBalance, topUpAmount, interestRate int) float32 {
	balance := startBalance + topUpAmount
	return float32(balance + startBalance*(interestRate/100))
}

func showSavedBalance(startBalance, topUpAmount, interestRate, numYears int) float32 {
	var accumulatedBalance float32
	for i := 1; i <= numYears; i++ {
		accumulatedBalance += calculateBalance(startBalance, topUpAmount, interestRate)
	}
	return accumulatedBalance
}

func main() {
	fmt.Println(showSavedBalance(1000, 200, 2, 2))
}
