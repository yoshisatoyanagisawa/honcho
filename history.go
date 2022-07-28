package main

import (
	"encoding/csv"
	"log"
	"os"
)

// loadHistory loads the history data from CSV and returns
// map from ID to History.
func loadHistory() map[string]History {
	file, err := os.Open("r3data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]History)
	for _, v := range rows[1:] {
		//     0,  1,     2,     3,           4,          5,
		// #kids, ID, grade, class, family name, first name,
		//    6,      7,     8,    9,   10
		// kana, region, phone, year, role
		if v[9] == "" && v[10] == "" {
			continue
		}
		h := History{
			ID: v[1],
			Year: v[9],
			Role: v[10],
			Phone: v[8],
		}
		m[h.ID] = h
	}
	return m
}
