package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"insectdie/kredit-pt-xyz/app"
	"insectdie/kredit-pt-xyz/controller"
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/middleware"
	"insectdie/kredit-pt-xyz/model/domain"
	"insectdie/kredit-pt-xyz/repository"
	"insectdie/kredit-pt-xyz/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/kredit_pt_xyz_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumenService := service.NewKonsumenService(konsumenRepository, db, validate)
	konsumenController := controller.NewKonsumenController(konsumenService)

	transaksiRepository := repository.NewTransaksiRepository()
	transaksiService := service.NewTransaksiService(transaksiRepository, db, validate)
	transaksiController := controller.NewTransaksiController(transaksiService)

	router := app.NewRouter(konsumenController, transaksiController)

	return middleware.NewAuthMiddleware(router)
}

func truncateKonsumen(db *sql.DB) {
	db.Exec("TRUNCATE konsumen")
}

func TestCreateKonsumenSuccess(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"nik" : "2222222222222222",
		"full_name" : "SANTI",
		"legal_name" : "SANTI",
		"tempat_lahir" : "PURWOKERTO",
		"tanggal_lahir" : "27/02/1999",
		"gaji" : 4500000,
		"foto_ktp" : "xyzasdzxsas.jpg",
		"foto_selfie" : "xyzasdzxsaasdasdsas.jpg"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/konsumens", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "SANTI", responseBody["data"].(map[string]interface{})["full_name"])
}

func TestCreateKonsumenFailed(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"nik" : "",
		"full_name" : "",
		"legal_name" : "",
		"tempat_lahir" : "",
		"tanggal_lahir" : "",
		"gaji" : 0,
		"foto_ktp" : "",
		"foto_selfie" : ""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/konsumens", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateKonsumenSuccess(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)

	tx, _ := db.Begin()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumen := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "2222222222222222",
		Full_name:     "SANTI",
		Legal_name:    "SANTI",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "27/02/1999",
		Gaji:          4500000,
		Foto_ktp:      "xyzasdzxsas.jpg",
		Foto_selfie:   "xyzasdzxsaasdasdsas.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"gaji" : 3000000, "foto_selfie" : "testupdate.jpg"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/konsumens/"+konsumen.Nik, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "testupdate.jpg", responseBody["data"].(map[string]interface{})["foto_selfie"])
}

func TestUpdateKonsumenFailed(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)

	tx, _ := db.Begin()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumen := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "2222222222222222",
		Full_name:     "SANTI",
		Legal_name:    "SANTI",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "27/02/1999",
		Gaji:          4500000,
		Foto_ktp:      "xyzasdzxsas.jpg",
		Foto_selfie:   "xyzasdzxsaasdasdsas.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"gaji" : 0, "foto_selfie" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/konsumens/"+konsumen.Nik, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetKonsumenSuccess(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)

	tx, _ := db.Begin()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumen := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "2222222222222222",
		Full_name:     "SANTI",
		Legal_name:    "SANTI",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "27/02/1999",
		Gaji:          4500000,
		Foto_ktp:      "xyzasdzxsas.jpg",
		Foto_selfie:   "xyzasdzxsaasdasdsas.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/konsumens/"+konsumen.Nik, nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, konsumen.Full_name, responseBody["data"].(map[string]interface{})["full_name"])
}

func TestGetKonsumenFailed(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/konsumens/404", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteKonsumenSuccess(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)

	tx, _ := db.Begin()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumen := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "2222222222222222",
		Full_name:     "SANTI",
		Legal_name:    "SANTI",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "27/02/1999",
		Gaji:          4500000,
		Foto_ktp:      "xyzasdzxsas.jpg",
		Foto_selfie:   "xyzasdzxsaasdasdsas.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/konsumens/"+konsumen.Nik, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteKonsumenFailed(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/konsumens/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListkonsumensSuccess(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)

	tx, _ := db.Begin()
	konsumenRepository := repository.NewKonsumenRepository()
	konsumen1 := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "1111111111111111",
		Full_name:     "BAMBANG",
		Legal_name:    "BAMBANG",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "10/10/2000",
		Gaji:          10000000,
		Foto_ktp:      "Qqweqwe.jpg",
		Foto_selfie:   "qwewqeewq.jpg",
	})
	konsumen2 := konsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:           "2222222222222222",
		Full_name:     "SANTI",
		Legal_name:    "SANTI",
		Tempat_lahir:  "PURWOKERTO",
		Tanggal_lahir: "27/02/1999",
		Gaji:          4500000,
		Foto_ktp:      "xyzasdzxsas.jpg",
		Foto_selfie:   "xyzasdzxsaasdasdsas.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/konsumens", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var konsumens = responseBody["data"].([]interface{})

	konsumenResponse1 := konsumens[0].(map[string]interface{})
	konsumenResponse2 := konsumens[1].(map[string]interface{})

	assert.Equal(t, konsumen1.Full_name, konsumenResponse1["full_name"])

	assert.Equal(t, konsumen2.Full_name, konsumenResponse2["full_name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateKonsumen(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/konsumens", nil)
	request.Header.Add("X-API-Key", "WRONG")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
