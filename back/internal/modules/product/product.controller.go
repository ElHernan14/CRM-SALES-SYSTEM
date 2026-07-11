package product

import (
	_ "crm-system-sales/docs"
	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	productdto "crm-system-sales/internal/modules/product/dto"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	service ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var req productdto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body:", err)
		return errorHandler.NewAppError(http.StatusBadRequest, "body inválido")
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.Create(r.Context(), &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query()
	req := productdto.GetProductsRequest{Search: q.Get("search"), Kind: q.Get("kind"), Type: q.Get("type"), Category: q.Get("category"), Page: 1, Limit: 10}
	if req.Kind == "" && (req.Type == "product" || req.Type == "service") {
		req.Kind = req.Type
		req.Type = ""
	}

	if v := q.Get("category_id"); v != "" {
		categoryID, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "category_id debe ser un nÃƒÆ’Ã‚Âºmero entero")
		}
		req.CategoryID = &categoryID
	}
	if v := q.Get("type_id"); v != "" {
		typeID, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "type_id debe ser un numero entero")
		}
		req.TypeID = &typeID
	}
	if v := q.Get("min_price"); v != "" {
		min, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "min_price debe ser un nÃƒÂºmero vÃƒÂ¡lido")
		}
		req.MinPrice = min
	}
	if v := q.Get("max_price"); v != "" {
		max, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "max_price debe ser un nÃƒÂºmero vÃƒÂ¡lido")
		}
		req.MaxPrice = max
	}
	if req.MinPrice > 0 && req.MaxPrice > 0 && req.MinPrice > req.MaxPrice {
		return errorHandler.NewAppError(http.StatusBadRequest, "min_price no puede ser mayor que max_price")
	}
	if v := q.Get("page"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "page debe ser un nÃƒÂºmero entero")
		}
		if p > 0 {
			req.Page = p
		}
	}
	if v := q.Get("limit"); v != "" {
		l, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "limit debe ser un nÃƒÂºmero entero")
		}
		if l > 0 {
			req.Limit = l
		}
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.GetProducts(r.Context(), &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id inválido")
	}

	res, err := c.service.GetByID(r.Context(), id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id inválido")
	}

	var req productdto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "body json inválido")
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.Update(r.Context(), id, &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *ProductController) UploadProductImage(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id inválido")
	}

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "imagen invalida")
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "image es requerido")
	}
	defer file.Close()

	res, err := c.service.UploadImage(r.Context(), id, file, header)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id inválido")
	}

	if err := c.service.Delete(r.Context(), id); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(nil))
}

func (c *ProductController) GetCompanyProducts(w http.ResponseWriter, r *http.Request) error {
	req, err := helper.ParseGetCompanyProductsRequest(r)
	if err != nil {
		return err
	}

	res, err := c.service.GetCompanyProducts(r.Context(), req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}
