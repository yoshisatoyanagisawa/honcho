package main

import (
	"sort"
)

// Sort kids in each family with kids' grade and class.
// Smaller grade/class should go first.
func sortKids(fs []Family) {
	for _, v := range fs {
		sort.SliceStable(v.Kids, func(i, j int) bool {
			if v.Kids[i].Grade == v.Kids[j].Grade {
				return v.Kids[i].Class < v.Kids[j].Class
			}
			return v.Kids[i].Grade < v.Kids[j].Grade
		})
	}
}
