package main

import "fmt"

type human interface {
	speak()
}

type agent struct {
	first string
}

func (a agent) speak(){
	fmt.Println("---")
}

func check(h human){
	h.speak()
	/*
		h.(type) when using with switch you can do a kind of polymorphism check
	 */
}


func sum(slice []int) int{
	ret:=0

	for _,v := range slice {
		ret+=v
	}
	return ret


}

func even (f func([]int) int, slice []int ) int {
	var filtered []int

	for _,v := range slice {
		if v%2==0{
			filtered=append(filtered,v)
		}
	}

	return f(filtered)
}

func tuple(i int) (int,string){
	return i*2,fmt.Sprintf("%v",i*16)
}

var g func (int) = func(a int){
	fmt.Println("------->>",a)
}

func makeEvenGenerator() func() int {
	v:=0
	return func () int{
		v+=2
		return v
	}
}

func main() {


	fmt.Println(tuple(2))
	g(3)

	nextEven:=makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())


}
