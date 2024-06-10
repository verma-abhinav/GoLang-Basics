package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new Bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0}
	return b
}

// function to format bill
func (thisBill *bill) formatbill() string {
	formatedString := "   Your Bill Breakdown : \n"
	var total float64 = 0
	// list of items
	for key, value := range thisBill.items {
		formatedString += fmt.Sprintf("%-25v $%v \n", key+":", value)
		total += value
	}
	// total
	total += thisBill.tip
	formatedString += fmt.Sprintf("%-25v $%0.2f \n", "tip:", thisBill.tip)
	formatedString += fmt.Sprintf("%-25v $%0.2f \n", "total:", total)
	formatedString += fmt.Sprintf("%v visiting us !!! ", "    Thank you for")
	return formatedString
}

// function to update the tip
func (thisBill *bill) updateTip(tip float64) {
	thisBill.tip = tip

}

// function to Add item to the Bill
func (thisBill *bill) addItems(item string, cost float64) {
	thisBill.items[item] = cost

}
