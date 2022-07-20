package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func atoi(s string) int {
	r, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return r
}

func loadHistory() map[string]History {
	j, err := ioutil.ReadFile("history.json")
	if err != nil {
		log.Fatal(err)
	}

	var hs []History
	if err := json.Unmarshal(j, &hs); err != nil {
		log.Fatal(err)
	}
	m := make(map[string]History)
	for _, v := range hs {
		m[v.ID] = v
	}
	return m
}

func main() {
	file, err := os.Open("B2022.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	hm := loadHistory()

	fs := make(map[string]Family)
	for _, v := range rows[1:] {
		//     0,  1,     2,     3,           4,          5,
		// #kids, ID, grade, class, family name, first name,
		//    6,      7,     8,    9,   10
		// kana, region, phone, year, role
		ID := v[1]

		var h *History
		if vv, ok := hm[ID]; ok {
			h = &vv
		} else if v[9] != "" {
			h = &History{
				ID,
				v[9],
				v[10],
				v[8],
			}
		}

		if val, ok := fs[ID]; ok {
			val.Kids = append(val.Kids, Kid{
				v[5],
				atoi(v[2]),
				atoi(v[3]),
			})
			fs[ID] = val
		} else {
			fs[ID] = Family{
				ID,
				v[4],
				[]Kid{
					{
						v[5],
						atoi(v[2]),
						atoi(v[3]),
					},
				},
				v[8],
				h,
			}
		}
	}
	var fsDump []Family
	for _, v := range fs {
		fsDump = append(fsDump, v)
	}
	fmt.Println(fsDump)
	out, err := json.Marshal(&fsDump)
	if err != nil {
		log.Fatal(err)
	}
	fout, err := os.Create("out.json")
	if err != nil {
		log.Fatal(err)
	}
	fout.Write(out)
	err = fout.Close()
	if err != nil {
		log.Fatal(err)
	}
}
