package main

import (
	"fmt"
	"log"
	"strconv"
)

func atoi(s string) int {
	r, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return r
}

// loadHistory loads the history data from CSV and returns
// map from ID to History.
func loadHistory() (map[string]History, error) {
	rows, err := loadCSV("r3data.csv")
	if err != nil {
		return nil, err
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
			ID:    v[1],
			Year:  v[9],
			Role:  v[10],
			Phone: v[8],
		}
		m[h.ID] = h
	}
	return m, nil
}

func mergeHistory(h1, h2 *History) *History {
	if h1 == nil {
		return h2
	}
	if h1.ID != h2.ID || h1.Phone != h2.Phone {
		log.Fatalf("ID (%s, %s) or Phone (%s, %s)are different ",
			h1.ID, h2.ID, h1.Phone, h2.Phone)
	}
	h := &History{
		ID:    h1.ID,
		Year:  fmt.Sprintf("%s,%s", h1.Year, h2.Year),
		Role:  fmt.Sprintf("%s,%s", h1.Role, h2.Role),
		Phone: h1.Phone,
	}
	return h
}

func main() {
	rows, err := loadCSV("B2022.csv")
	if err != nil {
		log.Fatal(err)
	}

	hm, err := loadHistory()
	if err != nil {
		log.Fatal(err)
	}

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
		}
		if v[9] != "" {
			h = mergeHistory(h,
				&History{
					ID:    ID,
					Year:  v[9],
					Role:  v[10],
					Phone: v[8],
				})
		}

		if val, ok := fs[ID]; ok {
			val.Kids = append(val.Kids, Kid{
				FirstName: v[5],
				Furigana:  v[6],
				Grade:     atoi(v[2]),
				Class:     atoi(v[3]),
			})
			fs[ID] = val
		} else {
			fs[ID] = Family{
				ID:         ID,
				FamilyName: v[4],
				Kids: []Kid{
					{
						FirstName: v[5],
						Furigana:  v[6],
						Grade:     atoi(v[2]),
						Class:     atoi(v[3]),
					},
				},
				Phone:   v[8],
				Region:  v[7],
				History: h,
			}
		}
	}
	var fsDump []Family
	for _, v := range fs {
		fsDump = append(fsDump, v)
	}
	fmt.Println(fsDump)
	err = storeJSON(fsDump, "out.json")
	if err != nil {
		log.Fatal(err)
	}
}
