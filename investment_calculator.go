package main

import "math"
import "fmt"

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	exceptedReturnRate := 5.5
	var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+exceptedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureRealValue)
}
