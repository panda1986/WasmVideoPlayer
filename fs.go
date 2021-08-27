package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	port := ":8888"
	log.Println(fmt.Sprintf("listen on %v", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}