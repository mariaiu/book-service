package repository

import (
	"context"
	"github.com/mariaiu/book/internal/models"
)

//go:generate mockgen -destination=mocks/repository.go -source=repository.go

type Repository interface {
	GetBooksByAuthor(ctx context.Context, author string) ([]models.Book, error)
	GetAuthorsByBook(ctx context.Context, book string) ([]models.Author, error)
}
