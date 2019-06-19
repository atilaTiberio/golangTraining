package main

import "fmt"

func main() {
	/*
	   scope of x only inside if block
	*/

	if x := 42; x == 42 {
		fmt.Println("----")
	}
/*
in this switch 2 cases matches for keeping flow execution just add fallthrough
 */
	/*switch {
	case false:
		fmt.Println("False")
	case (2 == 2):
		fmt.Println("2 igual a 2")
		fallthrough // everythings below it's ok even if it doesn't
	case (3 == 3):
		fmt.Println("3 igual a 3")
	default:
		fmt.Println("default")
	}*/

	s:= "other"

	switch s{
	case "Some","or","other":
		fmt.Println("3 possible matches")
	case "new":
		fmt.Println("new")
	default:
		fmt.Println("default")
	}


}
