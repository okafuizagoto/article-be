package article

import (
	jaegerLog "article-be/pkg/log"

	"github.com/opentracing/opentracing-go"
)

type IarticleSvc interface {
}

type Handler struct {
	articleSvc IarticleSvc
	tracer     opentracing.Tracer
	logger     jaegerLog.Factory
}

// New for bridging product handler initialization
func New(is IarticleSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		articleSvc: is,
		tracer:     tracer,
		logger:     logger,
	}
}
