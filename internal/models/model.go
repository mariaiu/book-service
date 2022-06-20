package models

type Book struct {
	ID    int    `json:"id,omitempty"    db:"id"`
	Title string `json:"title,omitempty" db:"title"`
}

type Author struct {
	ID   int    `json:"id,omitempty"   db:"id"`
	Name string `json:"name,omitempty" db:"name"`
}
