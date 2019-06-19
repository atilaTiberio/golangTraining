package main

import "fmt"

type person struct {
	first string
	last  string
}

/*func (p person) String() string{
	return fmt.Sprintf("Person[First: %v Last %v]",p.first,p.last)
}
*/
type secretAgent struct {
	person
	ltk bool
}

type NewAgent secretAgent
func main() {

	//VALUES OF TYPE

	//simpleStruct()
	//embeddedStruct()

	//anonymousStruct()

	mapOfStructs()


}

func mapOfStructs(){

	var mapita = map[string]person{
		"last_name":person{first:"f",last:"f"},
	}
	fmt.Println(mapita)

}

func anonymousStruct() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "first",
		last:  "last",
		age:   20,
	}

	fmt.Println(p1)
}

func embeddedStruct() {

	sa := secretAgent{person: person{first: "some", last: "some"}, ltk: true}
	fmt.Println(sa)
	fmt.Println(sa.first, sa.last, sa.ltk)

}

func simpleStruct() {
	//Value of type person
	p1 := person{first: "Some",
		last: "last",
	}

	slice := []person{p1}

	fmt.Println(slice)

	fmt.Println(p1.first, p1.last)
}
