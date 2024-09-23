package main

import (
	"backend/createListItem"
	"backend/getListItems"
	"fmt"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setupRoutes() {
	mux := http.NewServeMux() // function for setting up API routes

	// API ROUTE FOR CREATING LIST ITEM
	mux.HandleFunc("/createListItem", createListItem.CreateListItem)

	// API ROUTE FOR GETING ALL ITEMS VIA PROCEDURE
	mux.HandleFunc("/getListItems", getListItems.GetAllItems)
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
func main() {

	fmt.Println("Server is running on 8080 port")
	setupRoutes()
}
