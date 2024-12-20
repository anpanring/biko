package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Image string `json:"image"`
}

func readData(gender string) []Item {
	var items []Item

	file, err := os.Open(gender + "-products.csv")
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}

	for _, record := range records {
		items = append(items, Item{ID: record[1], Name: record[0], Link: record[2], Image: record[3]})
	}

	return items
}

func mensHandler(w http.ResponseWriter, r *http.Request) {
	items := readData("men")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func womensHandler(w http.ResponseWriter, r *http.Request) {
	items := readData("women")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/clothes/men", mensHandler)
	http.HandleFunc("/api/clothes/women", womensHandler)
	// log.Fatal("listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
