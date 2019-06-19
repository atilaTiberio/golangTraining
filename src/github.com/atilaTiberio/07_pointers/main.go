package main


import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	//c.radius=5

	c.radius=6
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5} //POINTER RECEIVER DOESN'T WORK WITH NON-POINTER VALUES
	/*c := new(circle)// with new you can get a pointer without a need to use * or &
	c.radius=5

	 */

	// c:=&circle{radius:5} it's much better to use new


	info(&c)
}