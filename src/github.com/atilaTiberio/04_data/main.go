package main

import (
	"fmt"
	"sort"
)

func main(){

/*
	Create, append, delete(append with boundaries)
 */
	//slices()
	//makeSlices()
	// multiDimSlices()
	//mapExample()
	//mapOfSlices()
	dataExercises()
}

func dataExercises(){

}

func mapOfSlices(){

	mapSlices := map[string][]string{
		"a":[]string{"a","b"},
		"b":[]string {"c","d"},
	}
	fmt.Println(mapSlices)
}

func mapExample(){

	m:= map[string]int {
		"a":1,
		"b":2,
		"c":20,
	}

	/*
		To check if a key exist
	 */
	if _,ok:=m["c"]; ok{
		fmt.Println("C key exists")
	}

	//delete(m,"a")
	m["d"]=30

	fmt.Println(m)
	for key, value := range m {
		fmt.Print("Key: ",key," value: ", value,"\n")
	}
	fmt.Println()

}


func multiDimSlices(){

	a := [] string{"a","b","c"}
	b := [] string{"d","e","f"}

	both := [][]string {a,b}
	fmt.Println(both)

	newBoth:= [][]string {[]string{"1","2","3"},[]string{"4","5","6"}}

	fmt.Println(newBoth)
}

func makeSlices(){

	/*
		commplex but valid
		var p *[]= new ([]int)
	*P = make([]int,100,100)

	Idiomatica

	when capacity is reached go automatically doubles the size
	 */
	funnySlice := make([]int,10,50)
	funnySlice=append(funnySlice,20)
	fmt.Println(funnySlice)

}

func slices() {
	/*
		Composite literal
	*/

	// var c:= type{values}
	/*
		Slices
	*/
	x:= [] int {3,4,5}


	y := [] int{12,45,87,89,21,43,65}

	fmt.Println(y[:7])

	fmt.Println(cap(y))
	for index, value := range y {
		fmt.Println(index,value)
	}

	/*
		Slicing a slice
	*/

	// inclusive:exclusive
	fmt.Println("Portion: ",y[1:3])
	fmt.Println("From 0 To limit defined: ",y[:3])
	fmt.Println("From 1 to the end:",y[1:])

	y = append(y, 44, 55, 66)
	fmt.Println(y)

	newSlice:= append(x,y...)
	fmt.Println("New Slice",newSlice)
	// There's no a delete function, so we should use then append func
	newSlice = append(newSlice[:2],newSlice[4:]...)

	fmt.Println("Result of deleting",newSlice)

	sort.Ints(newSlice)

	fmt.Println(newSlice)

}

