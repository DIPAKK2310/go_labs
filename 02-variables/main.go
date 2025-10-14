package main
import "fmt"

func main(){
	var userName = "Dipak Khare"
	fmt.Println(userName)
	//%T format specifier to print type of variable
	fmt.Printf("Data type of variable is %T \n " ,userName)
	
	var isLoggedIn = true
	fmt.Println(isLoggedIn)
	//%T format specifier to print type of variable
	fmt.Printf("Data type of variable is  %T \n" ,isLoggedIn)

	//var smallFloat = 32.78656 is written by shorthand ":=" it detect data type also only in function not on pakcge level
	 smallFloat:= 32.78656
	fmt.Println(smallFloat)
	//%T format specifier to print type of variable
	fmt.Printf("Data type of variable is: %T \n" ,smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("data type of variable is: %T \n", anotherVariable)
}
