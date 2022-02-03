package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var text string
	if option1 < option2 {
		text = option1
	} else {
		text = option2
	}
	output := text + " is clearly the better choice."
	return output
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discount float64
	switch {
	case age <= 3:
		discount = 0.8
	case age > 3 && age < 10:
		discount = 0.7
	case age >= 10:
		discount = 0.5
	}
	output := originalPrice * discount
	return output
}
