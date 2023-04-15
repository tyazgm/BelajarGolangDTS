package repository

import (
	"ch2session5/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) Get(bookID ...string) ([]model.Book, error) {
	books := make([]model.Book, 0)

	tx := br.db.Find(&books, bookID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return books, nil
}

func (br *BookRepository) Save(newBook model.Book) (model.Book, error) {
	tx := br.db.Create(&newBook)

	if tx.Error != nil {
		return model.Book{}, tx.Error
	}

	return newBook, nil
}

func (br *BookRepository) Update(bookUpdate model.Book, bookID string) (model.Book, error) {
	tx := br.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "book_id",
			},
		}}).
		Where("book_id = ?", bookID).
		Updates(&bookUpdate)

	if tx.Error != nil {
		return model.Book{}, tx.Error
	}

	return bookUpdate, nil
}

func (br *BookRepository) Delete(bookID string) (model.Book, error) {
	var deletedBook model.Book

	tx := br.db.Clauses(clause.Returning{}).Delete(&deletedBook, bookID)

	if tx.Error != nil {
		return model.Book{}, tx.Error
	}

	return deletedBook, nil
}
