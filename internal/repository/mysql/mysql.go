package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mariaiu/book/internal/config"
	"github.com/mariaiu/book/internal/models"
)

type repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*repository, error) {
	mysqlPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		cfg.DB.Username, "mysql", cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)

	db, err := sql.Open("mysql", mysqlPath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%s", mysqlPath),
	)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return &repository{db: db}, nil
}

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) GetBooksByAuthor(ctx context.Context, author string) ([]models.Book, error) {
	var books []models.Book

	query := fmt.Sprintf(`SELECT b.id, b.title FROM book b 
                                  JOIN book_author ba ON b.id = ba.book_id 
								  JOIN author a ON a.id = ba.author_id
									WHERE  a.name = ?`)

	rows, err := r.db.QueryContext(ctx, query, author); if err != nil {
		return nil, err
	}

	defer rows.Close()


	for rows.Next() {
		var b models.Book
		if err = rows.Scan(&b.ID, &b.Title); err != nil {
			return nil, err
		}
		books = append(books, b)
	}


	return books, nil
}


func (r *repository) GetAuthorsByBook(ctx context.Context, book string) ([]models.Author, error) {
	var authors []models.Author

	query := fmt.Sprintf(`SELECT a.id, a.name FROM author a 
                                  JOIN book_author ba ON a.id = ba.author_id 
								  JOIN book b ON ba.book_id=b.id
									WHERE b.title=?`)

	rows, err := r.db.QueryContext(ctx, query, book); if err != nil {
			return nil, err
		}

		defer rows.Close()


	for rows.Next() {
		var a models.Author
		if err = rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}


	return authors, nil
}

