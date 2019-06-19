package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	s:= "passwoord"
	bs, err:= bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err !=nil {
		return
	}
	fmt.Println(string(bs))


}
