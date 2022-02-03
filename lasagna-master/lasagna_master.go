package lasagna

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minute int) (preparationTime int) {
	if minute == 0 {
		minute = 2
	}
	preparationTime = len(layers) * minute
	return preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var (
		sauceCount   int
		noodlesCount int
	)
	for _, v := range layers {
		if v == "noodles" {
			noodlesCount = noodlesCount + 1
		}
		if v == "sauce" {
			sauceCount = sauceCount + 1
		}
	}
	noodles = noodlesCount * 50
	sauce = float64(sauceCount) * 0.2
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) (myNewList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (output []float64) {
	var data float64
	amounts := float64(portions) / 2.0
	fmt.Println(quantities)
	for _, v := range quantities {
		data = v * amounts
		output = append(output, data)
	}
	return output
}
