package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const morningPatrolDays = 6

func loadFinished() map[string]bool {
	rows, err := loadCSV("2022finished.csv")
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]bool)
	for _, v := range rows[2:] {
		//    0,      1,    2,     3,     4
		// date, circle, name, phone, grade
		m[v[3]] = true
	}
	return m
}

func filterDone(fs []Family, phone map[string]bool, done_id map[string]bool) []Family {
	var newfs []Family
	for _, v := range fs {
		if phone[v.Phone] {
			continue
		}
		if done_id[v.ID] {
			continue
		}
		newfs = append(newfs, v)
	}
	return newfs
}

func swapRow(rs [][]string, i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
	rs[i][0], rs[j][0] = rs[j][0], rs[i][0]
}

func toCSV(dates []time.Time, fs []Family, isCircle map[string]bool) [][]string {
	var rs [][]string
	var circleIdx []int
	var roleIdx []int
	dateIdx := 0
	for i, v := range fs {
		var firsts []string
		var grades []string
		for _, k := range v.Kids {
			firsts = append(firsts, k.FirstName)
			grades = append(grades, fmt.Sprintf("%d", k.Grade))
		}
		name := fmt.Sprintf("%s %s", v.FamilyName, strings.Join(firsts, "/"))
		date := ""
		circle := ""
		if isCircle[v.ID] {
			date = dates[dateIdx].Format("2006-01-02")
			dateIdx++
			circle = "〇"
			circleIdx = append(circleIdx, i)
			if v.History != nil && v.History.Year == "R4" {
				roleIdx = append(roleIdx, i)
			}
		}
		role := ""
		if v.History != nil && v.History.Year == "R4" {
			role = "R4"
		}
		rs = append(rs, []string{
			// date, circle, name, phone, grade
			date,
			circle,
			name,
			strings.Join(grades, ""),
			v.Phone,
			role,
		})
	}
	if len(roleIdx) >= 1 {
		swapRow(rs, roleIdx[0], circleIdx[0])
	}
	if len(roleIdx) >= 2 {
		swapRow(rs, roleIdx[1], circleIdx[len(circleIdx)-1])
	}
	return rs
}

func updateDone(m map[string]bool, s string) {
	if m[s] {
		log.Fatal("duplicated ID:" + s)
	}
	m[s] = true
}

func nextDate(t time.Time) time.Time {
	t = t.AddDate(0, 0, 1)
	for t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

func isInHoliday(t time.Time) bool {
	winterBegin, err := time.Parse("2006-01-02", "2022-12-24")
	if err != nil {
		log.Fatal(err)
	}
	winterEnd, err := time.Parse("2006-01-02", "2023-01-07")
	if err != nil {
		log.Fatal(err)
	}
	springBegin, err := time.Parse("2006-01-02", "2023-03-26")
	if err != nil {
		log.Fatal(err)
	}
	if t.After(winterBegin) && t.Before(winterEnd) {
		return true
	}
	if t.After(springBegin) {
		log.Fatal("we cannot allocate time: ", t)
	}
	return false
}

func nextWeek(t time.Time) time.Time {
	t = t.AddDate(0, 0, 7)
	for isInHoliday(t) {
		t = t.AddDate(0, 0, 7)
	}
	return t
}

func verifyPhoneExists(fs []Family, phone map[string]bool) {
	pcnt := 0
	for _, f := range fs {
		if phone[f.Phone] {
			fmt.Println(f)
			pcnt++
		}
	}
	fmt.Println("#Apr patrol", pcnt)
}

func main() {
	fs, err := loadJSON("out.json")
	if err != nil {
		log.Fatal(err)
	}

	sortFamilyWithGrade(fs, false)
	pm := loadFinished()
	verifyPhoneExists(fs, pm)
	done := make(map[string]bool) // ID -> bool (true if done)

	// choose morning patrol in Oct
	infs := filterDone(fs, pm, done)
	outfs := []Family{}
	isCircle := make(map[string]bool)
	dates := []time.Time{}
	date, err := time.Parse("2006-01-02", "2022-10-03")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < morningPatrolDays; i++ {
		dates = append(dates, date)
		date = nextDate(date)
		// with circle
		circle := infs[i]
		outfs = append(outfs, circle)
		isCircle[circle.ID] = true
		updateDone(done, circle.ID)
		// no circle
		buddy := infs[len(infs)*2/3+i]
		outfs = append(outfs, buddy)
		updateDone(done, buddy.ID)
	}
	storeCSV(toCSV(dates, outfs, isCircle), "morning_oct.csv")
	fmt.Println("done morning oct")

	// choose afternoon patrol
	infs = filterDone(fs, pm, done)
	outfs = []Family{}
	isCircle = make(map[string]bool)
	dates = []time.Time{}
	date, err = time.Parse("2006-01-02", "2022-08-29")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(infs))
	bSize := len(infs) / 3
	for i := 0; i < bSize; i++ {
		dates = append(dates, date)
		date = nextWeek(date)
		// with circle
		circle := infs[i]
		outfs = append(outfs, circle)
		isCircle[circle.ID] = true
		updateDone(done, circle.ID)
		// no circle
		buddy := infs[len(infs)*2/3-i-1]
		outfs = append(outfs, buddy)
		updateDone(done, buddy.ID)
		// no circle 2
		buddy = infs[len(infs)-i-1]
		outfs = append(outfs, buddy)
		updateDone(done, buddy.ID)
	}
	infs = filterDone(fs, pm, done)
	fmt.Println(infs)
	outfs = append(outfs, infs...)
	storeCSV(toCSV(dates, outfs, isCircle), "afternoon.csv")
	fmt.Println("done afternoon")
}
