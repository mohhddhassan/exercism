// Package weather provides tool to forecast the weather from city and condition.
package weather

var (
    //CurrentCondition variable declared as string to be used in the function Forecast.
	CurrentCondition string
    //CurrentLocation variable declared as string to be used in the function Forecast.
	CurrentLocation  string
)

// Forecast function returns the current location with it current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
