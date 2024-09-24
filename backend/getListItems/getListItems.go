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
	// Connect to the database
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
	titleQuery := r.URL.Query().Get("title")
	idQuery := r.URL.Query().Get("id")

	var rows *sql.Rows

	// Prepare the SQL query based on provided parameters
	if idQuery != "" {
		// If ID is provided, search by ID
		stmt, err := database.Prepare("SELECT * FROM get_all_items() WHERE id = $1 ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing statement: %v", err), http.StatusInternalServerError)
			return
		}
		defer func(stmt *sql.Stmt) {
			err := stmt.Close()
			if err != nil {

			}
		}(stmt)
		rows, err = stmt.Query(idQuery)
	} else if titleQuery != "" {
		// If title is provided, search by title
		stmt, err := database.Prepare("SELECT * FROM get_all_items() WHERE title ILIKE '%' || $1 || '%' ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing statement: %v", err), http.StatusInternalServerError)
			return
		}
		defer func(stmt *sql.Stmt) {
			err := stmt.Close()
			if err != nil {

			}
		}(stmt)
		rows, err = stmt.Query(titleQuery)
	} else {
		// Fetch all items if no search query
		rows, err = database.Query("SELECT * FROM get_all_items() ORDER BY created_at DESC")
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error calling procedure: %v", err), http.StatusInternalServerError)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows) // Ensure rows are closed after use

	var items []map[string]interface{}

	// Iterate through the rows
	for rows.Next() {
		var id, price, quantity int
		var title, description, created_at, updated_at string

		// Scan the row into variables
		err := rows.Scan(&id, &title, &description, &price, &quantity, &created_at, &updated_at)
		if err != nil {
			http.Error(w, "Error scanning the row", http.StatusInternalServerError)
			return
		}

		// Create a map for the current item
		item := map[string]interface{}{
			"id":          id,
			"title":       title,
			"description": description,
			"price":       price,
			"quantity":    quantity,
			"created_at":  created_at,
			"updated_at":  updated_at,
		}

		// Append the item to the items slice
		items = append(items, item)
	}

	// Check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error with rows", http.StatusInternalServerError)
		return
	}

	// Check if items slice is empty and return an explicit empty array
	if len(items) == 0 {
		response := map[string]interface{}{
			"items": []interface{}{}, // Explicitly return an empty array
		}
		// Set response header and encode items to JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
		return
	}

	// Prepare the response map
	response := map[string]interface{}{
		"items": items,
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
