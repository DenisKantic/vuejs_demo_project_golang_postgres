package main

import (
	"backend/createUpdateListItem"
	"backend/deleteListItem"
	"backend/getListItems"
	"fmt"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	mux.HandleFunc("/createListItem", createUpdateListItem.CreateListItem)
	// API ROUTE FOR UPDATING ONE ITEM
	mux.HandleFunc("/updateListItem", createUpdateListItem.UpdateListItem)

	// API ROUTE FOR GETTING ALL ITEMS VIA PROCEDURE
	mux.HandleFunc("/getListItems", getListItems.GetAllItems)
	// API ROUTE FOR GETTING ONE ITEM VIA ID QUERY
	mux.HandleFunc("/getOneItem", getListItems.GetOneItem)

	// API ROUTE FOR DELETING ITEM VIA ID QUERY PARAMETER
	mux.HandleFunc("/deleteItem", deleteListItem.DeleteItem)
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
func main() {

	fmt.Println("Server is running on 8080 port")
	setupRoutes()
}
