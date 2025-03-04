package main

import ("fmt")

func bmi_calculator() {
	var weight, height float64

	fmt.Print("Please enter your weight in kg: ")
	fmt.Scan(&weight)

	fmt.Print("Please enter your height in meter: ")
	fmt.Scan(&height)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Health Status: underweight")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Health Status: Normal Weight")
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Health Status: Overweight")
	} else {
		fmt.Println("Health Status: Obese")
	}

}

func main() {
	bmi_calculator()
}