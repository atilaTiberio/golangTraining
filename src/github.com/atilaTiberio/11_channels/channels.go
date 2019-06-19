package main

import (
	"fmt"
)

func main(){
	// doesNotRun()
	// doesRun()
	// bufferedChannels()
	sendAndReceiveChannels()

}

func sendAndReceiveChannels(){

	c := make(chan int )

	//send
	go send(c)
	receive(c)

}
//send
func send(c chan <- int){

	c<- 42
}

//receive

func receive(c <- chan int){
	fmt.Println(<-c)
}


func bufferedChannels(){
	c := make(chan int,2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func doesRun()  {
	c := make (chan int)

	go func(){
		 c <- 42
	}()
	fmt.Println(<-c)
}

func doesNotRun(){
	c := make(chan int)
	//c := make(chan int, 1) with this will work

	c <- 42

	fmt.Println(<-c)
}