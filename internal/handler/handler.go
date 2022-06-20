package handler

import (
	"context"
	"errors"
	"github.com/mariaiu/book/internal/repository"
	pb "github.com/mariaiu/book/proto"
)

type Handler struct {
	pb.UnimplementedBookServer
	repo    repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{
		repo:   repo,
		UnimplementedBookServer: pb.UnimplementedBookServer{},
	}
}
var (
	ErrEmptyRequest = errors.New("can't be empty")
)

func (h *Handler) GetBooksByAuthor(ctx context.Context, r *pb.GetBooksByAuthorRequest) (*pb.GetBooksByAuthorResponse, error) {
	if r.GetAuthor() == "" {
		return nil, ErrEmptyRequest
	}

	books, err := h.repo.GetBooksByAuthor(ctx, r.GetAuthor()); if err != nil {
		return nil, err
	}

	booksResp :=  make([]*pb.GetBooksByAuthorResponse_Book, 0, len(books))

	for _, book := range books {
		booksResp = append(booksResp, &pb.GetBooksByAuthorResponse_Book{
			Title: book.Title,
		})
	}

	return &pb.GetBooksByAuthorResponse{Books: booksResp}, nil
}

func (h *Handler) GetAuthorsByBook(ctx context.Context, r *pb.GetAuthorsByBookRequest) (*pb.GetAuthorsByBookResponse, error) {
	if r.GetBook() == "" {
		return nil, ErrEmptyRequest
	}

	authors, err := h.repo.GetAuthorsByBook(ctx, r.GetBook()); if err != nil {
		return nil,  err
	}

	authorsResp :=  make([]*pb.GetAuthorsByBookResponse_Author, 0, len(authors))

	for _, author := range authors {
		authorsResp = append(authorsResp, &pb.GetAuthorsByBookResponse_Author{
			Name: author.Name,
		})
	}

	return &pb.GetAuthorsByBookResponse{Author: authorsResp}, nil
}


