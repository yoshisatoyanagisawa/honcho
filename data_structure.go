package main

// History represents each ID's history.
type History struct {
	ID    string `json:"ID"`
	Year  string `json:"year"`
	Role  string `json:"role"`
	Phone string `json:"phone"` // used only for data verification
}

// Kid represents each kid information.
type Kid struct {
	FirstName string `json:"first name"`
	Furigana  string `json:"furigana"`
	Grade     int    `json:"grade"`
	Class     int    `json:"class"`
}

// Family represents each family information.
type Family struct {
	ID         string   `json:"ID"`
	FamilyName string   `json:"family name"`
	Kids       []Kid    `json:"kids"`
	Phone      string   `json:"phone"`
	Region     string   `json:"region"`
	History    *History `json:"history"`
}
