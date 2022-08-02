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

func sortKidsRev(fs []Family) {
	for _, v := range fs {
		sort.SliceStable(v.Kids, func(i, j int) bool {
			if v.Kids[i].Grade == v.Kids[j].Grade {
				return v.Kids[i].Class > v.Kids[j].Class
			}
			return v.Kids[i].Grade > v.Kids[j].Grade
		})
	}
}

func sortFamilyWithGrade(fs []Family, smallKidsFirst bool) {
	if smallKidsFirst {
		sortKids(fs)
	} else {
		sortKidsRev(fs)
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
}
