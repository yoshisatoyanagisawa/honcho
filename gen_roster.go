/*
 * This is a program to generate a list of addresses used for labels.
 * The generated CSV file will be used with a label maker software.
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
)

func main() {
	var (
		input  = flag.String("input", "input.json", "input file in JSON")
		output = flag.String("output", "output.csv", "output roster fiel in CSV")
	)
	flag.Parse()
	fs, err := loadJSON(*input)
	if err != nil {
		log.Fatal(err)
	}
	sortKidsRev(fs)

	// Sort families with the first kid grade and class.
	// Smaller grade/class should go first.
	// Using a family name as tie breaker.
	sort.SliceStable(fs, func(i, j int) bool {
		return fs[i].ID < fs[j].ID
	})
	fmt.Println(fs)
	var rs [][]string // output records
	for _, v := range fs {
		for _, k := range v.Kids {
			h := &History{}
			if v.History != nil {
				h = v.History
			}
			rs = append(rs, []string{
				// num kids, ID, grade, class, family name,
				// first name, phone, history
				fmt.Sprintf("%d", len(v.Kids)),
				v.ID,
				fmt.Sprintf("%d", k.Grade),
				fmt.Sprintf("%d", k.Class),
				v.FamilyName,
				k.FirstName,
				k.Furigana,
				v.Region,
				v.Phone,
				h.Year,
				h.Role,
			})
		}
	}
	err = storeCSV(rs, *output)
	if err != nil {
		log.Fatal(err)
	}
}
