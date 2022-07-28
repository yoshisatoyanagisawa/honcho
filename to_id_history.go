package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
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

	var hs []History
	for _, v := range rows[1:] {
		if v[9] == "" && v[10] == "" {
			continue
		}
		//     0,  1,     2,     3,           4,          5,
		// #kids, ID, grade, class, family name, first name,
		//    6,      7,     8,    9,   10
		// kana, region, phone, year, role
		hs = append(hs, History{
			v[1],
			v[9],
			v[10],
			v[8],
		})
	}
	fmt.Println(hs)
	out, err := json.Marshal(&hs)
	if err != nil {
		log.Fatal(err)
	}
	fout, err := os.Create("history.json")
	if err != nil {
		log.Fatal(err)
	}
	fout.Write(out)
	err = fout.Close()
	if err != nil {
		log.Fatal(err)
	}
}
