package adapter

import (
	"github.com/julienschmidt/httprouter"
	"golang11_restful_api/common"
	"golang11_restful_api/domain/entity"
	"golang11_restful_api/usecase"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	uc usecase.CategoryUsecase
}

func NewCategoryHandler(uc usecase.CategoryUsecase) CategoryController {
	return &CategoryHandler{uc}
}

func (h CategoryHandler) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	input := &usecase.CategoryInput{}
	common.GetRequestBody(request, &input)

	output := h.uc.Create(request.Context(), input)

	response := &common.PublicResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	common.SendResponseBody(writer, &response)
}

func (h CategoryHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	common.PanicIfError(err)

	input := &usecase.CategoryInput{Id: id}
	common.GetRequestBody(request, &input)

	output := h.uc.Update(request.Context(), input)

	response := &common.PublicResponse{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	common.SendResponseBody(writer, &response)
}

func (h CategoryHandler) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	common.PanicIfError(err)

	category := &entity.Category{
		Id: id,
	}

	h.uc.Delete(request.Context(), id)

	response := &common.PublicResponse{
		Code:   http.StatusOK,
		Status: "Deleted",
		Data:   usecase.NewCategoryOutput(category),
	}

	writer.Header().Add("Content-Type", "application/json")

	common.SendResponseBody(writer, &response)
}

func (h CategoryHandler) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	common.PanicIfError(err)

	output := h.uc.FindById(request.Context(), id)

	response := &common.PublicResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	common.SendResponseBody(writer, &response)
}

func (h CategoryHandler) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	output := h.uc.FindAll(request.Context())

	response := &common.PublicResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	common.SendResponseBody(writer, &response)
}
