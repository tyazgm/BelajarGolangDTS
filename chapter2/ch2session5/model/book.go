package model

type Book struct {
	BookID   string `gorm:"primary key" json:"book_id"`
	Title    string `gorm:"not null;unique;varchar(255)" json:"title"`
	Price    int    `gorm:"not null" json:"price"`
	AuthorID int    `json:"author_id"`
}
