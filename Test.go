package main

import (
	"fmt"
	"strconv"
)

func Test() {
	var TestData = make(map[string]string)
	var TestData1 = make(map[string]string)
	var TestData2 = make(map[string]string)
	var TestData3 = make(map[string]string)

	// Same currency Test
	TestData["sourceCurrency"] = "NGN"
	TestData["destinationCurrency"] = "NGN"
	TestData["amount"] = "3000"
	TestData["output"] = "3000"

	// Different currency Test
	TestData1["sourceCurrency"] = "NGN"
	TestData1["destinationCurrency"] = "KES"
	TestData1["amount"] = "3000"
	TestData1["output"] = "823"

	// Negative result
	TestData2["sourceCurrency"] = "GHS"
	TestData2["destinationCurrency"] = "NGN"
	TestData2["amount"] = "-3000"
	TestData2["output"] = "-1"

	// Different currency Test
	TestData3["sourceCurrency"] = "KES"
	TestData3["destinationCurrency"] = "GHS"
	TestData3["amount"] = "3000"
	TestData3["output"] = "169"

	outcome := CurrencyConverter(TestData)
	o0, _ := strconv.ParseFloat(TestData["output"], 64)
	if o0 == outcome {
		fmt.Println("Test 1 Passed - Same currency conversion")
	} else {
		fmt.Println("Test 1 Failed - Same currency conversion")
	}

	outcome1 := CurrencyConverter(TestData1)
	o1, _ := strconv.ParseFloat(TestData1["output"], 64)
	if o1 == outcome1 {
		fmt.Println("Test 2 Passed - Cross-currency conversion")
	} else {
		fmt.Println("Test 2 Failed - Cross-currency conversion")
	}

	outcome2 := CurrencyConverter(TestData2)
	o2, _ := strconv.ParseFloat(TestData2["output"], 64)
	if o2 == outcome2 {
		fmt.Println("Test 3 Passed - Negative input conversion")
	} else {
		fmt.Println("Test 3 Failed - Negative input conversion")
	}

	outcome3 := CurrencyConverter(TestData2)
	o3, _ := strconv.ParseFloat(TestData2["output"], 64)
	if o3 == outcome3 {
		fmt.Println("Test 3 Passed - Cross-currency conversion")
	} else {
		fmt.Println("Test 3 Failed - Cross-currency conversion")
	}
}
