package article

import (
	httpHelper "article-be/internal/delivery/http"
	articleEntity "article-be/internal/entity/article"
	"article-be/pkg/response"
	"encoding/json"
	"io/ioutil"
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
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var (
		result        interface{}
		metadata      interface{}
		err           error
		resp          response.Response
		updateArticle articleEntity.Put
		types         string
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
	if id != "" {
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &updateArticle)
		ids, _ := strconv.Atoi(id)
		result, err = h.articleSvc.UpdateArticle(ctx, ids, updateArticle)
		if err != nil {
			log.Println("err", err)
		}
	}
	// Your code here
	types = r.FormValue("type")
	switch types {

	}

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
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
