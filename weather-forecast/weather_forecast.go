//Package weather is for package.
package weather

//CurrentCondition is for variable.
var CurrentCondition string

//CurrentLocation is for variable.
var CurrentLocation string

//Forecast is for function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
