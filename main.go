package main

import (
	"net/http"
	"log"
	"excel/controllers"
	"time"
	"fmt"
)

func main()  {

	//http.HandleFunc("/",controllers.Index)

	blogC := controllers.BlogController{}
	http.HandleFunc("/",blogC.Home)

	http.HandleFunc("/index",blogC.Index)

	go test2()

	//http://127.0.0.1:3000/public/
	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))
	log.Println("start web server :3000")
	log.Fatal(http.ListenAndServe(":3000",nil))

}

func test2()  {

	t := time.NewTicker(time.Second * 5)

	for v := range t.C{
		log.Println(v)
	}
}
func test1()  {

	t := time.NewTicker(time.Second * 5)

	for{
		select {
			case a :=<-t.C:
				fmt.Println(a)
		}
	}
}
