package main

type Content struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Image    string `json:"image"`
	Creator  string `json:"creator"`
	Content  string `json:"content"`
	Active   string `json:"active"`
	Link     string `json:"link"`
}

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Image    string `json:"image"`
	Link     string `json:"link"`
}