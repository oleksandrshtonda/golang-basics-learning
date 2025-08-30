package main

import "math"
import "fmt"

func main() {
	var investmentAmount = 1000
	var exceptedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+exceptedReturnRate/100, float64(years))

	fmt.Println(futureValue)
}
