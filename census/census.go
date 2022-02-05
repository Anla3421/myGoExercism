// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	output := &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return output
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	var (
		k string
		v string
	)

	for k, v = range r.Address {
	}
	if r.Name == "" || r.Age < 0 || k == "" || v == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	var nilMap map[string]string
	r.Name = ""
	r.Age = 0
	r.Address = nilMap
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var (
		ak     string
		av     string
		output int
	)

	for _, v := range residents {
		for ak, av = range v.Address {
			if v.Name != "" || v.Age != 0 || ak != "" || av != "" {
				output = output + 1
			}
		}
	}
	return output
}
