package main

import (
	"net/http"
	"log"
	"excel/controllers"
)

func main()  {

	//http.HandleFunc("/",controllers.Index)

	blogC := controllers.BlogController{}
	http.HandleFunc("/",blogC.Home)

	http.HandleFunc("/index",blogC.Index)

	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))
	log.Println("start web server :3000")
	log.Fatal(http.ListenAndServe(":3000",nil))

}
