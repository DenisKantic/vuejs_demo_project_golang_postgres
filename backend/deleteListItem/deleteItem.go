package deleteListItem

import (
	"backend/db_connection"
	"net/http"
	"strconv"
)

// DeleteItem handles the HTTP request to delete an item by ID using a query parameter
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Parse item ID from query parameters
	idStr := r.URL.Query().Get("id") // Get the 'id' query parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	// Connect to the database
	database, err := db_connection.DBConnect()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	// Prepare and execute the delete statement
	stmt, err := database.Prepare("DELETE FROM item WHERE id = $1")
	if err != nil {
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the delete statement
	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
