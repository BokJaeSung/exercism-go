package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRate64 := float64(productionRate)
	return productionRate64 * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	answer := CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0
	return int(answer)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	mock := carsCount / 10
	namuge := carsCount % 10
	return uint(mock*95000 + namuge*10000)
}
