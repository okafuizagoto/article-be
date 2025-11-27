package article

import (
	httpHelper "article-be/internal/delivery/http"
	"article-be/pkg/response"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// Getgoldgym godoc
// @Summary Get entries of all goldgyms
// @Description Get entries of all goldgyms
// @Tags goldgym
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200
// @Router /v1/profiles [get]
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getgoldgym", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	params := mux.Vars(r)
	id := params["id"]
	ids, _ := strconv.Atoi(id)
	if id != "" {
		result, err = h.articleSvc.DeleteArticle(ctx, ids)
	}
	types = r.FormValue("type")
	switch types {
	}
	value, ok := result.(string)
	if !ok {
		fmt.Println("not string")
	}
	if result != "Berhasil Delete" {
		resp = response.Response{
			Data:     result,
			Metadata: nil,
			Error: response.Error{
				Status: true,
				Msg:    value,
			},
			StatusCode: 400,
		}
	}
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		//
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	return
}
