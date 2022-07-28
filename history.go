package main

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
