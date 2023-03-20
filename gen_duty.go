package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	var (
		input      = flag.String("input", "input.json", "input file in JSON")
		candidates = flag.String("candidates", "candidates.csv", "A CSV file output that have candidates")
		done       = flag.String("done", "done.csv", "A CSV file output that have people who have already experienced the roles")
	)
	flag.Parse()
	fs, err := loadJSON(*input)
	if err != nil {
		log.Fatal(err)
	}

	sortFamilyWithGrade(fs, true)
	fmt.Println(fs)
	var rs [][]string // for whom has never had role before.
	ridx := 1
	var hrs [][]string // for whom has had a role before.
	hidx := 1
	for _, v := range fs {
		var firsts []string
		var grades []string
		for _, k := range v.Kids {
			firsts = append(firsts, k.FirstName)
			grades = append(grades, fmt.Sprintf("%d", k.Grade))
		}
		if v.History == nil {
			// We cannot ask this person to be the role.
			if len(grades) == 1 && grades[0] == "6" {
				fmt.Println("excluded:", v)
				continue
			}
			rs = append(rs, []string{
				// seq, familiy + list of first, grade
				fmt.Sprintf("%d", ridx),
				v.FamilyName,
				strings.Join(firsts, ","),
				strings.Join(grades, ""),
			})
			ridx++
		} else {
			hrs = append(hrs, []string{
				fmt.Sprintf("%d", hidx),
				v.FamilyName,
				strings.Join(firsts, ","),
				strings.Join(grades, ""),
				v.History.Year,
				v.History.Role,
			})
			hidx++
		}
	}

	err = storeCSV(rs, *candidates)
	if err != nil {
		log.Fatal(err)
	}

	err = storeCSV(hrs, *done)
	if err != nil {
		log.Fatal(err)
	}
}
