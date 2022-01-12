package adapter

import (
	"github.com/julienschmidt/httprouter"
	error2 "golang11_restful_api/common/error"
	"golang11_restful_api/common/handler"
	"golang11_restful_api/common/response"
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
	handler.GetRequestBody(request, &input)

	output := h.uc.Create(request.Context(), input)

	res := &response.PublicResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	handler.SendResponseBody(writer, &res, http.StatusCreated)
}

func (h CategoryHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	error2.PanicIfError(err)

	input := &usecase.CategoryInput{Id: id}
	handler.GetRequestBody(request, &input)

	output := h.uc.Update(request.Context(), input)

	res := &response.PublicResponse{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	handler.SendResponseBody(writer, &res, http.StatusOK)
}

func (h CategoryHandler) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	error2.PanicIfError(err)

	category := &entity.Category{
		Id: id,
	}

	h.uc.Delete(request.Context(), id)

	res := &response.PublicResponse{
		Code:   http.StatusOK,
		Status: "Deleted",
		Data:   usecase.NewCategoryOutput(category),
	}

	writer.Header().Add("Content-Type", "application/json")

	handler.SendResponseBody(writer, &res, http.StatusOK)
}

func (h CategoryHandler) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	error2.PanicIfError(err)

	output := h.uc.FindById(request.Context(), id)

	res := &response.PublicResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	handler.SendResponseBody(writer, &res, http.StatusOK)
}

func (h CategoryHandler) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	output := h.uc.FindAll(request.Context())

	res := &response.PublicResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   output,
	}

	writer.Header().Add("Content-Type", "application/json")

	handler.SendResponseBody(writer, &res, http.StatusOK)
}
