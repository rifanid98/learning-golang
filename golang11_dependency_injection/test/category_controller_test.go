package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"golang11_dependency_injection/adapter"
	error2 "golang11_dependency_injection/common/error"
	"golang11_dependency_injection/common/middleware"
	"golang11_dependency_injection/domain/entity"
	"golang11_dependency_injection/infrastructure/persistence/repository"
	router2 "golang11_dependency_injection/infrastructure/router"
	"golang11_dependency_injection/usecase"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func NewDBTest() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	error2.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryUsecase := usecase.NewCategoryInteractor(categoryRepository, db, validate)
	categoryController := adapter.NewCategoryHandler(categoryUsecase)

	router := router2.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Product"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "Created", responseBody["status"])
	assert.Equal(t, "Product", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Save(context.Background(), tx, &entity.Category{
		Name: "Product",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{"name": "Products"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Updated", responseBody["status"])
	assert.Equal(t, int(category.Id), int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Products", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Save(context.Background(), tx, &entity.Category{
		Name: "Product",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Save(context.Background(), tx, &entity.Category{
		Name: "Product",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, int(category.Id), int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Save(context.Background(), tx, &entity.Category{
		Name: "Product",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Deleted", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Save(context.Background(), tx, &entity.Category{
		Name: "Product",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	categories := responseBody["data"].([]interface{})
	category1 := categories[0].(map[string]interface{})

	assert.Equal(t, int(category.Id), int(category1["id"].(float64)))
	assert.Equal(t, category.Name, category1["name"])
}

func TestUnauthorized(t *testing.T) {
	db := NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
