package main

import "fmt"

func main() {
	fmt.Println("Main function")
	multiplesOfNine()
	//calling prime function
	if isPrime(4) {

		fmt.Println("given number is prime")
	} else {
		fmt.Println("given number is not prime")
	}
}

//Display name,id
func display(name string, id int) {
	fmt.Println(name, id)
}

//Display multiples of 9
func multiplesOfNine() {
	display("Manikanth Reddy", 500224276)
	fmt.Println("Multiples of 9")
	for i := 1; i < 10; i++ {
		fmt.Println(9 * i)
	}
}

//To check whether number is prime or not
func isPrime(number int) bool {
	display("Avinash", 500223817)
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}
