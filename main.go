package main

import (
	"fmt"
	"log"
	"net/http"
	"websockets-in-go/youtube"

	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println(" 📹 Youtube Subscriber Moniter")

	// Load env variables
	godotenv.Load(".env")

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", item)

	// setUpRoutes()
}
