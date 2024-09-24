package getListItems

import (
	"backend/db_connection"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"net/http"
	"strconv"
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

	// Get the search query from URL parameters
	query := r.URL.Query().Get("title")

	// Prepare the SQL query
	var rows *sql.Rows
	if query != "" {
		// Using a prepared statement to prevent SQL injection
		stmt, err := database.Prepare("SELECT * FROM get_all_items() WHERE title ILIKE '%' || $1 || '%' ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing statement: %v", err), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()
		rows, err = stmt.Query(query)
	} else {
		// Fetch all items if no search query
		rows, err = database.Query("SELECT * FROM get_all_items() ORDER BY created_at DESC")
	}
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

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	// Parse item ID from query parameters
	idStr := r.URL.Query().Get("id") // Get the 'id' query parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	fmt.Println("ID I GET", id)

	// Connect to the database
	database, err := db_connection.DBConnect()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)

	row := database.QueryRow("SELECT * FROM get_item_by_id($1)", id)

	// Handle potential errors
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching item", http.StatusInternalServerError)
		return
	}
	// Create an Item struct to hold the result
	var item Item
	err = row.Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Price,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	// Set response header to JSON and send the result
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

}
