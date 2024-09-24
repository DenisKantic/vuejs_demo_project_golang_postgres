package createUpdateListItem

import (
	"backend/db_connection"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

func CreateListItem(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB max request per form

	if err != nil {
		fmt.Println("Error on reading form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extracting form values
	title := r.FormValue("title")
	priceStr := r.FormValue("price")
	description := r.FormValue("description")
	quantityStr := r.FormValue("quantity")

	missingFields := []string{}

	if title == "" {
		missingFields = append(missingFields, "Title is empty")
	}
	if description == "" {
		missingFields = append(missingFields, "Description is empty")
	}
	// Convert quantity and price to integers
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		missingFields = append(missingFields, "Quantity must be greater than zero")
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil || price <= 0 {
		missingFields = append(missingFields, "Price must be greater than zero")
	}

	if len(missingFields) > 0 {
		// Return the error messages as JSON
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": missingFields,
		})
		return
	}

	// starting connection with database
	database, err := db_connection.DBConnect()
	if err != nil {
		fmt.Printf("Error connecting with database, %v\n", err)
		//http.Error(w, "Error has occured", http.StatusInternalServerError)
		return
	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			return
		}
	}(database)

	// calling the stored procedure
	_, err = database.Exec("CALL insert_item($1,$2,$3,$4)", title, description, price, quantity)
	if err != nil {
		fmt.Println("Error calling stored procedure")
		http.Error(w, "Failed executing form", http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "Item stored successfully")
}

func UpdateListItem(w http.ResponseWriter, r *http.Request) {
	// Parse the request to retrieve form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max request per form
	if err != nil {
		fmt.Println("Error on reading form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extracting form values
	idStr := r.FormValue("id") // Ensure that you have an ID field to identify the item to update
	title := r.FormValue("title")
	priceStr := r.FormValue("price")
	description := r.FormValue("description")
	quantityStr := r.FormValue("quantity")

	missingFields := []string{}

	// Validate required fields
	if title == "" {
		missingFields = append(missingFields, "Title is empty")
	}
	if description == "" {
		missingFields = append(missingFields, "Description is empty")
	}

	// Convert quantity and price to integers
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		missingFields = append(missingFields, "Quantity must be greater than zero")
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil || price <= 0 {
		missingFields = append(missingFields, "Price must be greater than zero")
	}

	// Ensure that ID is provided and valid
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		missingFields = append(missingFields, "Invalid item ID")
	}

	if len(missingFields) > 0 {
		// Return the error messages as JSON
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": missingFields,
		})
		return
	}

	// Connect to the database
	database, err := db_connection.DBConnect()
	if err != nil {
		fmt.Printf("Error connecting with database, %v\n", err)
		http.Error(w, "Error has occurred", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	// Call the stored procedure to update the item
	_, err = database.Exec("SELECT update_item($1, $2, $3, $4, $5)", id, title, description, price, quantity)
	if err != nil {
		fmt.Println("Error calling stored procedure")
		http.Error(w, "Failed executing form", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Item updated successfully")
}
