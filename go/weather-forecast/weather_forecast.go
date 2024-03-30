// Package weather provides tools to help you
// get information about the weather.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location, like a city name.
var CurrentLocation string

// Forecast returns a string informing the current weather in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
