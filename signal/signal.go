package main

import (
	"fmt"
	"os/signal"
	"os"
	"syscall"
)

func main()  {


	fmt.Println("start.....")
	ch := make(chan os.Signal)

	signal.Notify(ch,syscall.SIGINT)
	//signal.Notify(ch)

	
	for{
		sig := <-ch

		switch sig {

		case syscall.SIGINT:
			fmt.Println("sigint")
		case syscall.SIGTERM:
			fmt.Println("sigterm")
		default:
			fmt.Println(sig)

		}
		
	}
	
	

}
