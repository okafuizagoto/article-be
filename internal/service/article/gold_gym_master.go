package article

import (
	articleEntity "article-be/internal/entity/article"
	"article-be/pkg/errors"
	"context"
	"strings"
)

func (s Service) InsertArticle(ctx context.Context, article articleEntity.Post) (string, error) {
	var (
		message string
		err     error
	)
	validStatus := map[string]bool{
		"publish": true,
		"draft":   true,
		"trash":   true,
	}

	status := strings.ToLower(article.Status)

	if len(article.Title) > 20 {
		message = "Title melebihi 200 karakter"
		return message, err
	}
	if len(article.Content) > 200 {
		message = "Content melebihi 200 karakter"
		return message, err
	}
	if len(article.Category) > 3 {
		message = "Category melebihi 3 karakter"
		return message, err
	}

	if !validStatus[status] {
		message = "Status tidak sesuai (hanya bisa isi publish, draft, thrash)"
		return message, err
	}
	err = s.goldgym.InsertArticle(ctx, article)
	if err != nil {
		return message, errors.Wrap(err, "[Service][InsertArticle]")
	}
	message = "Berhasil Insert"

	return message, err
}

func (s Service) GetArticleByID(ctx context.Context, id string) (articleEntity.Get, error) {

	users, err := s.goldgym.GetArticleByID(ctx, id)
	if err != nil {
		return users, errors.Wrap(err, "[Service][GetGoldUser]")
	}
	return users, nil
}

func (s Service) GetArticleByPagination(ctx context.Context, page, length int) ([]articleEntity.Get, int, error) {
	var total int
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * length
	users, err := s.goldgym.GetArticleByPagination(ctx, offset, length)
	if err != nil {
		return users, total, errors.Wrap(err, "[Service][GetArticleByPagination]")
	}

	total, err = s.goldgym.GetAllArticleByPagination(ctx)
	if err != nil {
		return users, total, errors.Wrap(err, "[Service][GetArticleByPagination]")
	}
	return users, total, nil
}

func (s Service) UpdateArticle(ctx context.Context, id int, article articleEntity.Put) (string, error) {
	var (
		message string
		err     error
	)
	validStatus := map[string]bool{
		"publish": true,
		"draft":   true,
		"trash":   true,
	}

	status := strings.ToLower(article.Status)

	if len(article.Title) > 20 {
		message = "Title melebihi 200 karakter"
		return message, err
	}
	if len(article.Content) > 200 {
		message = "Content melebihi 200 karakter"
		return message, err
	}
	if len(article.Category) > 3 {
		message = "Category melebihi 3 karakter"
		return message, err
	}

	if !validStatus[status] {
		message = "Status tidak sesuai (hanya bisa isi publish, draft, thrash)"
		return message, err
	}
	err = s.goldgym.UpdateArticle(ctx, id, article)
	if err != nil {
		return message, errors.Wrap(err, "[Service][UpdateArticle]")
	}
	message = "Berhasil Update"

	return message, err
}

func (s Service) DeleteArticle(ctx context.Context, id int) (string, error) {
	var (
		message string
		err     error
	)
	err = s.goldgym.DeleteArticle(ctx, id)
	if err != nil {
		message = "Gagal Delete"
		return message, errors.Wrap(err, "[Service][DeleteArticle]")
	}
	message = "Berhasil Delete"

	return message, err
}
