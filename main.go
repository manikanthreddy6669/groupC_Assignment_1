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

	Celsius()

	sum, difference := calculate(21, 13)
	fmt.Println("Sum:", sum, "Difference:", difference)

	fmt.Printf("\n%d! = %d\n", 5, factorial(5))

	//Calculate square of numbers
	Sq(2)
	Sq(9)
	Sq(13)
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

func Celsius() {
	display("Shashank", 500224639)

	var ftemp float64 = 0
	var ctemp float64 = 0

	fmt.Printf("Enter temperature in Celsius: ")
	fmt.Scanf("%f", &ctemp)
	ftemp = (ctemp * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit: %.2f", ftemp)
}

func calculate(n1 int, n2 int) (int, int) {
	display("Nishith", 500223382)
	sum := n1 + n2
	difference := n1 - n2
	return sum, difference
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

//Calculate square of numbers
func Sq(num int) {
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}
