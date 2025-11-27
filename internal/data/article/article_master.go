package article

import (
	articleEntity "article-be/internal/entity/article"
	"article-be/pkg/errors"
	"context"
	"strconv"
	"strings"
)

func (d Data) InsertArticle(ctx context.Context, article articleEntity.Post) error {

	var err error

	_, err = (*d.stmt)[insertArticle].ExecContext(ctx,
		article.Title,
		article.Content,
		article.Category,
		article.Status,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) GetArticleByID(ctx context.Context, id string) (articleEntity.Get, error) {
	var (
		// article  articleEntity.Get
		articles articleEntity.Get
		err      error
	)
	ids, _ := strconv.Atoi(id)
	rows, err := (*d.stmt)[getArticleByID].QueryxContext(ctx, ids)
	if err != nil {
		return articles, errors.Wrap(err, "[DATA] [GetGoldUser]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&articles); err != nil {
			return articles, errors.Wrap(err, "[DATA] [GetGoldUser]")
		}
		// articles = append(articles, article)
	}
	return articles, err
}

func (d Data) GetArticleByPagination(ctx context.Context, offset, limit int) ([]articleEntity.Get, error) {
	var (
		article  articleEntity.Get
		articles []articleEntity.Get
		err      error
	)
	rows, err := (*d.stmt)[getArticleByPagination].QueryxContext(ctx, offset, limit)
	if err != nil {
		return articles, errors.Wrap(err, "[DATA] [GetArticleByPagination]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&article); err != nil {
			return articles, errors.Wrap(err, "[DATA] [GetArticleByPagination]")
		}
		articles = append(articles, article)
	}
	return articles, err
}

func (d Data) GetAllArticleByPagination(ctx context.Context) (int, error) {
	var (
		total int
		err   error
	)
	rows, err := (*d.stmt)[getAllArticleByPagination].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetAllArticleByPagination]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetAllArticleByPagination]")
		}
	}
	return total, err
}

func (d Data) UpdateArticle(ctx context.Context, id int, article articleEntity.Put) error {
	var (
		err error

		titlePtr, contentPtr, categoryPtr, statusPtr interface{}
	)

	if article.Title != "" {
		titlePtr = article.Title
	} else {
		titlePtr = nil
	}

	if article.Content != "" {
		contentPtr = article.Content
	} else {
		contentPtr = nil
	}

	if article.Category != "" {
		categoryPtr = article.Category
	} else {
		categoryPtr = nil
	}

	if article.Status != "" {
		statusPtr = strings.ToLower(article.Status)
	} else {
		statusPtr = nil
	}

	_, err = (*d.stmt)[updateArticle].ExecContext(ctx,
		titlePtr, contentPtr, categoryPtr, statusPtr,
		id)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateArticle]")
	}

	return err

}

func (d Data) DeleteArticle(ctx context.Context, id int) error {
	var (
		err error
	)

	_, err = (*d.stmt)[deleteArticle].ExecContext(ctx, id)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteArticle]")
	}

	return err

}
