package main

import (
	"fmt"
	"runtime"
)

type data struct {
	name string
}


func main() {

	s:= []data{{"one"}}
	fmt.Println(s)
	done := false

	go func() {
		done = true
	}()

	//for   {
	//	fmt.Println("bbbb")
	//}


	for !done{

		runtime.Gosched()

		fmt.Println("aaaa")
	}

	fmt.Println("!done")


}
