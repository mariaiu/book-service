package handler_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mariaiu/book/internal/handler"
	"github.com/mariaiu/book/internal/models"
	mock "github.com/mariaiu/book/internal/repository/mocks"
	pb "github.com/mariaiu/book/proto"
	"reflect"
	"testing"
)

func TestHandler_GetBooksByAuthor(t *testing.T) {
	ctx := context.Background()

	errInternal := errors.New("internal error")

	testCases := []struct {
		name       string
		request    *pb.GetBooksByAuthorRequest
		mockRes    []models.Book
		mockErr    error
		expectResp *pb.GetBooksByAuthorResponse
		expectErr  error
	}{
		{
			name:       "empty author's name",
			request:    &pb.GetBooksByAuthorRequest{Author: ""},
			mockRes:    nil,
			mockErr:    nil,
			expectResp: nil,
			expectErr: handler.ErrEmptyRequest,
		},
		{
			name:       "internal error",
			request:    &pb.GetBooksByAuthorRequest{Author: "Name"},
			mockRes:    nil,
			mockErr:    errInternal,
			expectResp: nil,
			expectErr: errInternal,
		},
		{
			name:       "success",
			request:    &pb.GetBooksByAuthorRequest{Author: "Name"},
			mockRes:    []models.Book{{Title: "Title"}},
			mockErr:    nil,
			expectResp: &pb.GetBooksByAuthorResponse{Books: []*pb.GetBooksByAuthorResponse_Book{{Title: "Title"}}},
			expectErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockService := mock.NewMockRepository(ctrl)
			srv := handler.NewHandler(mockService)

			if tc.mockRes != nil || tc.mockErr != nil {
				mockService.EXPECT().
					GetBooksByAuthor(ctx, gomock.Any()).
					Return(tc.mockRes, tc.mockErr)
			}

			resp, err := srv.GetBooksByAuthor(ctx, tc.request)

			if tc.expectErr != err {
				t.Errorf("expect %v, got %v", tc.expectErr, err)
			}

			if !reflect.DeepEqual(tc.expectResp, resp) {
				t.Errorf("expect %v, got %v", tc.expectResp, resp)
			}
		})
	}
}

func TestHandler_GetAuthorsByBook(t *testing.T) {
	ctx := context.Background()

	errInternal := errors.New("internal error")

	testCases := []struct {
		name       string
		request    *pb.GetAuthorsByBookRequest
		mockRes    []models.Author
		mockErr    error
		expectResp *pb.GetAuthorsByBookResponse
		expectErr  error
	}{
		{
			name:       "empty book's title",
			request:    &pb.GetAuthorsByBookRequest{Book: ""},
			mockRes:    nil,
			mockErr:    nil,
			expectResp: nil,
			expectErr: handler.ErrEmptyRequest,
		},
		{
			name:       "internal error",
			request:    &pb.GetAuthorsByBookRequest{Book: "Title"},
			mockRes:    nil,
			mockErr:    errInternal,
			expectResp: nil,
			expectErr: errInternal,
		},
		{
			name:       "success",
			request:    &pb.GetAuthorsByBookRequest{Book: "Title"},
			mockRes:    []models.Author{{Name: "Name"}},
			mockErr:    nil,
			expectResp: &pb.GetAuthorsByBookResponse{Author: []*pb.GetAuthorsByBookResponse_Author{{Name: "Name"}}},
			expectErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockService := mock.NewMockRepository(ctrl)
			srv := handler.NewHandler(mockService)

			if tc.mockRes != nil || tc.mockErr != nil {
				mockService.EXPECT().
					GetAuthorsByBook(ctx, gomock.Any()).
					Return(tc.mockRes, tc.mockErr)
			}

			resp, err := srv.GetAuthorsByBook(ctx, tc.request)

			if tc.expectErr != err {
				t.Errorf("expect %v, got %v", tc.expectErr, err)
			}

			if !reflect.DeepEqual(tc.expectResp, resp) {
				t.Errorf("expect %v, got %v", tc.expectResp, resp)
			}
		})
	}
}
