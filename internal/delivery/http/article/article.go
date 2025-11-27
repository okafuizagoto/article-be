package article

import (
	articleEntity "article-be/internal/entity/article"
	jaegerLog "article-be/pkg/log"
	"context"

	"github.com/opentracing/opentracing-go"
)

type IarticleSvc interface {
	InsertArticle(ctx context.Context, article articleEntity.Post) (string, error)
	GetArticleByID(ctx context.Context, id string) (articleEntity.Get, error)
	GetArticleByPagination(ctx context.Context, page, length int) ([]articleEntity.Get, int, error)
	UpdateArticle(ctx context.Context, id int, article articleEntity.Put) (string, error)
	DeleteArticle(ctx context.Context, id int) (string, error)
}

type IarticleSvcStock interface {
}

type (
	// Handler ...
	Handler struct {
		articleSvc      IarticleSvc
		articleSvcStock IarticleSvcStock
		tracer          opentracing.Tracer
		logger          jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IarticleSvc, isst IarticleSvcStock, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		articleSvc:      is,
		articleSvcStock: isst,
		tracer:          tracer,
		logger:          logger,
	}
}
