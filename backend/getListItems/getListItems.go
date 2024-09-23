package getListItems

import (
	"backend/db_connection"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// GetAllItems handles the HTTP request to get all items
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	database, err := db_connection.DBConnect()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := database.Close(); err != nil {
			fmt.Printf("Error closing database: %v\n", err)
		}
	}()

	// Call the stored procedure
	rows, err := database.Query("SELECT * FROM get_all_items()")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error calling procedure: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Close rows after use

	var adoptPost []map[string]interface{}

	// Iterate through the rows
	for rows.Next() {
		var id, price, quantity int
		var title, description, created_at, updated_at string

		err := rows.Scan(&id, &title, &description, &price, &quantity, &created_at, &updated_at)
		if err != nil {
			http.Error(w, "Error scanning the row", http.StatusInternalServerError)
			return
		}

		// Create a map for the current item
		post := map[string]interface{}{
			"id":          id,
			"title":       title,
			"description": description,
			"price":       price,
			"quantity":    quantity,
			"created_at":  created_at,
			"updated_at":  updated_at,
		}

		adoptPost = append(adoptPost, post)
	}

	// Check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error with rows", http.StatusInternalServerError)
		return
	}

	// Prepare the response map
	response := map[string]interface{}{
		"items": adoptPost, // Change key to "items"
	}

	// Set response header and encode items to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}
