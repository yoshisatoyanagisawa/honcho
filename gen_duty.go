package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
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
	for _, v := range fs {
		sort.SliceStable(v.Kids, func(i, j int) bool {
			return v.Kids[i].Grade < v.Kids[j].Grade
		})
	}
	sort.SliceStable(fs, func(i, j int) bool {
		nki := len(fs[i].Kids)
		nkj := len(fs[j].Kids)
		n := nki
		if n > nkj {
			n = nkj
		}
		for k := 0; k < n; k++ {
			if fs[i].Kids[k].Grade == fs[j].Kids[k].Grade {
				continue
			}
			return fs[i].Kids[k].Grade > fs[j].Kids[k].Grade
		}
		if nki == nkj {
			return fs[i].FamilyName > fs[j].FamilyName
		}
		return nki < nkj
	})
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

	f, err := os.Create("duty.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	w.WriteAll(rs)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Create("noduty.csv")
	if err != nil {
		log.Fatal(err)
	}
	w = csv.NewWriter(f)
	w.WriteAll(hrs)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
