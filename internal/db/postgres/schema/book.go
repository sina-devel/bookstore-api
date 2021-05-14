package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"gorm.io/gorm"
)

type (
	Book struct {
		gorm.Model
		Name          string
		Description   string
		File          string
		SellerID      uint
		CategoryID    uint
		Comments      []Comment
		DownloadCount uint
		Pictures      []Picture
		Status        types.BookStatus
		Price         types.Price
	}
)

func (b *Book) ConvertModel() *models.Book {
	var comments []models.Comment
	for _, comment := range b.Comments {
		comments = append(comments, models.Comment{
			ID:          comment.ID,
			UserID:      comment.UserID,
			Text:        comment.Text,
			BookID:      comment.BookID,
			FullName:    comment.FullName,
			IsConfirmed: comment.IsConfirmed,
		})
	}

	var pictures []models.Picture
	for _, picture := range b.Pictures {
		pictures = append(pictures, models.Picture{
			ID:     picture.ID,
			Name:   picture.Name,
			Alt:    picture.Alt,
			BookID: picture.BookID,
		})
	}

	return &models.Book{
		ID:            b.ID,
		Name:          b.Name,
		Description:   b.Description,
		File:          b.File,
		SellerID:      b.SellerID,
		CategoryID:    b.CategoryID,
		Comments:      comments,
		DownloadCount: b.DownloadCount,
		Pictures:      pictures,
		Status:        b.Status,
		Price:         b.Price,
	}

}

func ConvertBook(book *models.Book) *Book {
	var comments []Comment
	for _, comment := range book.Comments {
		comments = append(comments, Comment{
			Model:       gorm.Model{ID: comment.ID},
			UserID:      comment.UserID,
			Text:        comment.Text,
			BookID:      comment.BookID,
			FullName:    comment.FullName,
			IsConfirmed: comment.IsConfirmed,
		})
	}

	var pictures []Picture
	for _, picture := range book.Pictures {
		pictures = append(pictures, Picture{
			Model:  gorm.Model{ID: picture.ID},
			Name:   picture.Name,
			Alt:    picture.Alt,
			BookID: picture.BookID,
		})
	}

	return &Book{
		Model:         gorm.Model{ID: book.ID},
		Name:          book.Name,
		Description:   book.Description,
		File:          book.File,
		SellerID:      book.SellerID,
		CategoryID:    book.CategoryID,
		Comments:      comments,
		DownloadCount: book.DownloadCount,
		Pictures:      pictures,
		Status:        book.Status,
		Price:         book.Price,
	}
}
