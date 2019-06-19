package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

/*
	If we define this with lower case, marshaller is not going to work, we need to export with upperCase
*/
type person struct {
	First string
	Last  string
	Age   int
}

type ByAge []person

func (byAge ByAge) Len() int{
	return len(byAge)
}

func (byAge ByAge) Less(i,j int) bool{
	return byAge[i].Age <byAge[j].Age
}

func (byAge ByAge) Swap(i,j int ){
	byAge[i],byAge[j] = byAge[j], byAge[i]
}

var p1 = person{
First: "James",
Last:  "Bond",
Age:   32,
}

var p2 = person{
First: "Miss",
Last:  "Moneypenny",
Age:   27,
}

func main() {

	jsonProcess()

	//sortSlices()
	//customSortSlices()

}

func customSortSlices() {

	byAge:=ByAge([]person{p1,p2})
	fmt.Println(ByAge(byAge))
	sort.Sort(byAge)
	fmt.Println(ByAge(byAge))
}

func sortSlices() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
}

func jsonProcess() {

	people := [] person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	var peopleUnmarshalle []person
	error := json.Unmarshal(bs, &peopleUnmarshalle)

	if error != nil {
		fmt.Println("Error")
	}

	fmt.Println("New, content", peopleUnmarshalle)

	fmt.Fprintln(os.Stdout, "Hey std")
	io.WriteString(os.Stdout, "Hey\n")

	err = json.NewEncoder(os.Stdout).Encode(people)

	if err != nil{
		fmt.Println("rrrr")
	}




}
