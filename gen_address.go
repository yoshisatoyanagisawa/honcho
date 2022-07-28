/*
 * This is a program to generate a list of addresses used for labels.
 * The generated CSV file will be used with a label maker software.
 */
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	j, err := ioutil.ReadFile("out.json")
	if err != nil {
		log.Fatal(err)
	}

	var fs []Family
	if err := json.Unmarshal(j, &fs); err != nil {
		log.Fatal(err)
	}
	// Sort kids in each family with kids' grade and class.
	// Smaller grade/class should go first.
	for _, v := range fs {
		sort.SliceStable(v.Kids, func(i, j int) bool {
			if v.Kids[i].Grade == v.Kids[j].Grade {
				return v.Kids[i].Class < v.Kids[j].Class
			}
			return v.Kids[i].Grade < v.Kids[j].Grade
		})
	}
	// Sort families with the first kid grade and class.
	// Smaller grade/class should go first.
	// Using a family name as tie breaker.
	sort.SliceStable(fs, func(i, j int) bool {
		if fs[i].Kids[0].Grade != fs[j].Kids[0].Grade {
			return fs[i].Kids[0].Grade < fs[j].Kids[0].Grade
		}
		if fs[i].Kids[0].Class != fs[j].Kids[0].Class {
			return fs[i].Kids[0].Class < fs[j].Kids[0].Class
		}
		if fs[i].FamilyName != fs[j].FamilyName {
			return fs[i].FamilyName < fs[j].FamilyName
		}
		return fs[i].Kids[0].FirstName < fs[j].Kids[0].FirstName
	})
	fmt.Println(fs)
	var rs [][]string // output records
	for i, v := range fs {
		rs = append(rs, []string{
			// seq, familiy, first kids first, grade, class
			fmt.Sprintf("%d", i),
			v.FamilyName,
			v.Kids[0].FirstName,
			fmt.Sprintf("%d", v.Kids[0].Grade),
			fmt.Sprintf("%d", v.Kids[0].Class),
		})
	}

	f, err := os.Create("address.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	w.WriteAll(rs)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
