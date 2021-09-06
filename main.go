package main

import (
	"fmt"
	"log"
	"net/http"
	"websockets-in-go/websocket"

	"github.com/joho/godotenv"
)

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println(" 📹 Youtube Subscriber Moniter")

	// Load env variables
	godotenv.Load(".env")
	setUpRoutes()
	// item, err := youtube.GetSubscribers()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%+v\n", item)

	// setUpRoutes()
}
