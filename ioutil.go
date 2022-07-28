/*
 * I/O utility library
 */
package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func loadJSON(filename string) ([]Family, error) {
	j, err := os.ReadFile("out.json")
	if err != nil {
		return nil, err
	}

	var fs []Family
	if err := json.Unmarshal(j, &fs); err != nil {
		return nil, err
	}
	return fs, nil
}

func writeJSONFile(fs []Family, filename string) error {
	out, err := json.Marshal(&fs)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, out, os.ModePerm)
}

func loadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	return r.ReadAll()
}

func writeCSVFile(records [][]string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	w.WriteAll(records)
	return f.Close()
}
