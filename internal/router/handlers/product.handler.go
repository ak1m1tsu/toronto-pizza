package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

type ProductHandler struct {
	svc service.IProductService
}

func NewProductHandler(svc service.IProductService) *ProductHandler {
	return &ProductHandler{svc}
}

func (h *ProductHandler) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusOK, map[string]any{"message": "added"})
	JSON(w, resp.Status, resp)
}

func (h *ProductHandler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusOK, map[string]any{"message": "updated"})
	JSON(w, resp.Status, resp)
}

func (h *ProductHandler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusOK, map[string]any{"message": "deleted"})
	JSON(w, resp.Status, resp)
}

func (h *ProductHandler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusOK, map[string]any{"message": "got"})
	ctx := r.Context()

	product, ok := ctx.Value(ProductKey).(*models.ProductDTO)
	if !ok {
		resp.Status = http.StatusUnprocessableEntity
		resp.Body = map[string]any{"error": "uprocessable entity"}
		JSON(w, resp.Status, resp)
		return
	}

	resp.Body = map[string]any{"product": product}
	JSON(w, resp.Status, resp)
}

func (h *ProductHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusOK, map[string]any{"message": "got all"})
	JSON(w, resp.Status, resp)
}

func (h *ProductHandler) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := NewApiResponse(http.StatusNotFound, map[string]any{})
		productID := chi.URLParam(r, "id")
		product, err := h.svc.GetProductByID(r.Context(), productID)
		if err != nil {
			resp.Body = map[string]any{"error": "product not found"}
			JSON(w, resp.Status, resp)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
