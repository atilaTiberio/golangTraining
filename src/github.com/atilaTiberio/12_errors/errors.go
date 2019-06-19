package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){

	errorHandling()
}

func errorHandling() {

	f, err := os.Open("test.txt")

	if err!=nil {
		log.Println(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err !=nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))


}
