package main


import "fmt"

func main() {
	fmt.Println("Main function")
	multiplesOfNine()
}

//Display name,id
func display(name string, id int) {
    fmt.Println(name,id)
}

//Display multiples of 9
func multiplesOfNine() {
	display("Manikanth Reddy",500224276)
	fmt.Println("Multiples of 9")
	for i:=1;i<10;i++ {
	fmt.Println(9*i)
	}
}