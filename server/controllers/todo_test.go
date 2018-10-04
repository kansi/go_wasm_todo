package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kansi/go_wasm_todo/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ControllersTestSuite struct {
	suite.Suite
	DB     *models.Db
	router *gin.Engine
}

func (suite *ControllersTestSuite) SetupSuite() {
	db := models.NewDbConnection("/tmp/go_wasm_todo_test.db")
	suite.DB = db
	suite.router = Router(db)
}

func (suite *ControllersTestSuite) TearDownSuite() {
	suite.DB.Conn.Close()
}

// SetupTest clear database before each test
func (suite *ControllersTestSuite) SetupTest() {
	iter := suite.DB.Conn.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		suite.DB.Conn.Delete(key, nil)
	}
}

func (suite *ControllersTestSuite) TestGetTodo404() {
	req, _ := http.NewRequest("GET", "/todo/1", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)
	assert.Equal(suite.T(), 404, resp.Code)
}

func (suite *ControllersTestSuite) TestGetTodo() {
	todo := seedTodo(suite.DB)
	req, _ := http.NewRequest("GET", "/todo/some_random_id", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	expTodo := models.Todo{}
	_ = json.Unmarshal(resp.Body.Bytes(), &expTodo)

	assert.Equal(suite.T(), *todo, expTodo)
	assert.Equal(suite.T(), 200, resp.Code)
}

func (suite *ControllersTestSuite) TestDeleteTodo() {
	_ = seedTodo(suite.DB)
	req, _ := http.NewRequest("DELETE", "/todo/some_random_id", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), 200, resp.Code)
}

func TestTodoControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllersTestSuite))
}

// Store random todo
func seedTodo(db *models.Db) *models.Todo {
	stringDate := "2018-09-16T17:53:18+02:00"
	date, _ := time.Parse(time.RFC3339, stringDate)
	todo := models.Todo{
		ID:    "some_random_id",
		Title: "Random title",
		Body:  "Random body",
		Date:  date,
	}
	db.StoreTodo(&todo)

	return &todo
}
