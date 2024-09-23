package createListItem

import (
	"fmt"
	"net/http"
)

func createListItem(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB max request per form

	if err != nil {
		fmt.Println("Error on reading form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extracting form values
	name := r.FormValue("name")
	price := r.FormValue("price")
	description := r.FormValue("description")
	
}
