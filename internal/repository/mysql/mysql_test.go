package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mariaiu/book/internal/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetBooksByAuthor(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	var b models.Book

	query := fmt.Sprintf(`SELECT b.id, b.title FROM book b 
                                  JOIN book_author ba ON b.id = ba.book_id 
								  JOIN author a ON a.id = ba.author_id
									WHERE  a.name = ?`)

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(b.ID, b.Title)

	mock.ExpectQuery(query).WithArgs("name").WillReturnRows(rows)

	user, err := repo.GetBooksByAuthor(context.Background(),"name")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestGetAuthorsByBook(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	var a models.Author

	query := fmt.Sprintf(`SELECT a.id, a.name FROM author a 
                                  JOIN book_author ba ON a.id = ba.author_id 
								  JOIN book b ON ba.book_id=b.id
									WHERE b.title=?`)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(a.ID, a.Name)

	mock.ExpectQuery(query).WithArgs("name").WillReturnRows(rows)

	user, err := repo.GetAuthorsByBook(context.Background(),"name")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
