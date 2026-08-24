package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	c:= successRate/100
    wc:= c * float64(productionRate)
    return wc
    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	c:= CalculateWorkingCarsPerHour(productionRate,successRate)
    res := c/60
    r := int(res)
    return r 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    c := carsCount/10
    cc := carsCount%10
	cost := uint((c*95000) + (cc*10000))
    return cost
}
