package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	createProductDTO := &models.CreateProductDTO{}
	if err := json.NewDecoder(r.Body).Decode(&createProductDTO); err != nil {
		JSON(w, http.StatusBadRequest, nil, "", err)
		return
	}

	if err := models.ValidateCreateProductDTO(createProductDTO); err != nil {
		JSON(w, http.StatusBadRequest, nil, "", err)
		return
	}

	productDto, err := h.svc.InsertProduct(r.Context(), createProductDTO)
	if err != nil {
		JSON(w, http.StatusBadRequest, nil, "", err)
	}

	data := map[string]interface{}{"product": productDto}
	JSON(w, http.StatusOK, data, "Successfuly added.", nil)
}

func (h *ProductHandler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(keyProduct).(*models.ProductDTO)
	updateProduct := &models.UpdateProductDTO{}
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		JSON(w, http.StatusBadRequest, nil, "", err)
		return
	}
	product, err := h.svc.UpdateProduct(r.Context(), product.ID, updateProduct)
	if err != nil {
		JSON(w, http.StatusBadRequest, nil, "", err)
		return
	}

	data := map[string]interface{}{"updated": product}
	JSON(w, http.StatusOK, data, "Successfuly updated.", nil)
}

func (h *ProductHandler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(keyProduct).(*models.ProductDTO)
	id, err := h.svc.DeleteProduct(r.Context(), product.ID)
	if err != nil {
		JSON(w, http.StatusBadGateway, nil, "", err)
		return
	}
	data := map[string]interface{}{"deleted": id}
	JSON(w, http.StatusOK, data, "Successfuly deleted.", nil)
}

func (h *ProductHandler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	var (
		product = r.Context().Value(keyProduct).(*models.ProductDTO)
		data    = map[string]interface{}{"product": product}
	)
	JSON(w, http.StatusOK, data, "Successfuly received.", nil)
}

func (h *ProductHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	prodFilter := models.NewProductFilter(r)
	prodSort := []*models.ProductSort{models.NewProductSort(r)}
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil || page < 1 {
		page = 1
	}
	products, err := h.svc.GetProducts(r.Context(), prodFilter, prodSort, page)
	if err != nil {
		JSON(w, http.StatusInternalServerError, nil, "", err)
		return
	}
	data := map[string]interface{}{"products": products}
	JSON(w, http.StatusOK, data, "Successfuly received.", nil)
}

func (h *ProductHandler) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "id")
		product, err := h.svc.GetProductByID(r.Context(), productID)
		if err != nil {
			JSON(w, http.StatusNotFound, nil, "", ErrProductNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), keyProduct, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
