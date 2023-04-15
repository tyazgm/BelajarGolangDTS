package model

type Author struct {
	ID    int    `gorm:"primary key" json:"author_id"`
	Name  string `gorm:"not null;varchar(255)" json:"author_name"`
	Books []Book `json:"author_books"`
}
