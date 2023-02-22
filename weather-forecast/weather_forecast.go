// Package weather provides tools to return 
// the weather conditions in a specified location.
package weather

// CurrentCondition is a string describing the current weather condition.
var CurrentCondition string
// CurrentLocation is the location where the user wishes to know the weather.
var CurrentLocation string

// Forecast returns a sentence describing the weather in a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
