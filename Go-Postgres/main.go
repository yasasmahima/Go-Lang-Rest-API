package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
	 "os"
	 "github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	r := Router.Router()
	var port =  os.Getenv("PORT")
	fmt.Println("App starts on the port 8080...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
