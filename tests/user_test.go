package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
	"github.com/grootkng/clean-arch-golang/internal/pkg/factory"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var USER_ID int = 0

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestShouldReturn428PassingNoQueryParameters(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path, factory.UserFactory().FindAll)
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, string("{\"message\":\"Send page correctly\"}"), string(responseData))
	assert.Equal(t, http.StatusPreconditionRequired, w.Code)
}

func TestShouldReturn428PassingOnlyPageAsQueryParameter(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path, factory.UserFactory().FindAll)
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	q := req.URL.Query()
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, string("{\"message\":\"Send pageSize correctly\"}"), string(responseData))
	assert.Equal(t, http.StatusPreconditionRequired, w.Code)
}

func TestShouldReturn428PassingOnlyPageSizeAsQueryParameter(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path, factory.UserFactory().FindAll)
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	q := req.URL.Query()
	q.Add("pageSize", "1")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, string("{\"message\":\"Send page correctly\"}"), string(responseData))
	assert.Equal(t, http.StatusPreconditionRequired, w.Code)
}

func TestFindAllShouldReturn404WhenNoRecord(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path, factory.UserFactory().FindAll)

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("pageSize", "1")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, string("{\"message\":\"record not found\"}"), string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateUser(t *testing.T) {
	payload := `{"name":"Alice","gender":"Female","age":30}`
	path := "/v1/users"

	r := SetUpRouter()
	r.POST(path, factory.UserFactory().Create)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFindAllUsersWithContent(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path, factory.UserFactory().FindAll)

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("pageSize", "1")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	u := []entity.User{}
	response, _ := io.ReadAll(w.Body)
	if err := json.Unmarshal(response, &u); err != nil {
		t.Fail()
	}

	assert.Equal(t, 1, len(u))
	assert.Equal(t, http.StatusOK, w.Code)

	USER_ID = u[0].Id
}

func TestGetUserById(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.GET(path+"/:id", factory.UserFactory().FindBy)
	req, _ := http.NewRequest(http.MethodGet, path+"/"+strconv.Itoa(USER_ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	payload := `{"name":"Alice","gender":"Female","age":35}`
	path := "/v1/users"

	r := SetUpRouter()
	r.PUT(path+"/:id", factory.UserFactory().UpdateBy)
	req, _ := http.NewRequest(http.MethodPut, path+"/"+strconv.Itoa(USER_ID), bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	path := "/v1/users"

	r := SetUpRouter()
	r.DELETE(path+"/:id", factory.UserFactory().DeleteBy)
	req, _ := http.NewRequest(http.MethodDelete, path+"/"+strconv.Itoa(USER_ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
