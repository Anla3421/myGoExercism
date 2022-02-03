package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	output := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return output
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	output := make(map[string]int)
	return output
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitCheck := Units()
	_, exist := unitCheck[unit]
	if exist != true {
		return false
	}
	bill[item] = bill[item] + units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExist := bill[item]
	if itemExist != true {
		return false
	}
	unitCheck := Units()
	_, unitExist := unitCheck[unit]
	if unitExist != true {
		return false
	}
	switch {
	case bill[item] < units[unit]:
		return false
	case bill[item] == units[unit]:
		delete(bill, item)
	default:
		bill[item] = bill[item] - units[unit]
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, itemExist := bill[item]
	if itemExist != true {
		return 0, false
	}
	return quantity, true
}
