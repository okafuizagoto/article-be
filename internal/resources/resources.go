package resources

import (
	// sesuaikan path

	"article-be/internal/service/article"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"
)

type BootResources struct {
	DBLocal         *sqlx.DB
	DBProd          *sqlx.DB
	Redis           *redis.Client
	ArticleSvcLocal article.Service
	ArticleSvcProd  article.Service
	Tracer          opentracing.Tracer
	Logger          *zap.Logger
}
