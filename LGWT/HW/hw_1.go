package main

import "fmt"

const russia = "President"
const germany = "Chancellor"
const britain = "Queen"
const txt = " is the "

func main() {
	var country_input, name_input string
	fmt.Println("Enter country (1 out of 3: Russia/Germany/Britain) and the rulers name")
	fmt.Scan(&country_input, &name_input)
	fmt.Println(ruller(country_input, name_input))

}

func ruller(country, name string) (result string) {
	switch country {
	case "Russia":
		result = name + txt + russia
	case "Germany":
		result = name + txt + germany
	case "Britain":
		result = name + txt + britain
	default:
		result = "Incorrect input for country name"
	}
	return result
}
