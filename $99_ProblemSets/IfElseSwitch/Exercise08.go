package main

import "fmt"

const (
	FIRST_KM_FEE = 15000
	NEXT_KM_FEE  = 13500
	SIXTH_KM_FEE = 11000
)

func totalTaxiFare(distance int) float64 {
	if distance < 2 {
		return FIRST_KM_FEE * float64(distance)
	}
	if distance < 6 {
		return FIRST_KM_FEE + NEXT_KM_FEE*float64(distance-1)
	}
	if distance < 120 {
		return FIRST_KM_FEE + NEXT_KM_FEE*4 + SIXTH_KM_FEE*float64(distance-5)
	}
	return 0.9 * (FIRST_KM_FEE + NEXT_KM_FEE*4 + SIXTH_KM_FEE*float64(distance-5))
}

func main() {
	fmt.Print("Enter the distance (km) you have traveled:")
	var distance int
	fmt.Scan(&distance)

	totalFare := totalTaxiFare(distance)
	fmt.Printf("Total taxi fare you have to pay: $%.2f", totalFare)
}
