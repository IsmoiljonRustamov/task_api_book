package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createBook(t *testing.T) *Book {
	book, err := dbMangaer.Create(&Book{
		Title:      "Good person",
		AuthorName: "Ismoiljon",
		Price:      1222.00,
		Amount:     2,
	})
	require.NoError(t, err)
	require.NotEmpty(t, book)

	return book
}

func deleteBook(id int64, t *testing.T) {
	err := dbMangaer.Delete(id)
	require.NoError(t, err)
}

func TestCreate(t *testing.T) {
	b := createBook(t)
	deleteBook(b.Id, t)
}

func TestUpdate(t *testing.T) {
	b := createBook(t)

	b.Title = "Bad person"
	b.AuthorName = "Kimdir"
	b.Price = 299.00
	b.Amount = 1

	book, err := dbMangaer.Update(b)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	deleteBook(book.Id, t)
}
