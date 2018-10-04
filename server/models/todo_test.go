package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
	"time"
)

type ModelsTestSuite struct {
	suite.Suite
	DBConn *Db
}

func (suite *ModelsTestSuite) SetupSuite() {
	db, err := leveldb.OpenFile("/tmp/go_wasm_todo_test.db", nil)
	if err != nil {
		panic(fmt.Sprintf("Unable to open db file"))
	}
	suite.DBConn = &Db{Conn: db}
}

func (suite *ModelsTestSuite) TearDownSuite() {
	suite.DBConn.Conn.Close()
}

func (suite *ModelsTestSuite) TestDbInterface() {
	stringDate := "2018-09-16T17:53:18+02:00"
	date, _ := time.Parse(time.RFC3339, stringDate)

	todo := &Todo{
		ID:    "some_random_id",
		Title: "Random title",
		Body:  "Random body",
		Date:  date,
	}

	err := suite.DBConn.StoreTodo(todo)
	assert.Nil(suite.T(), err)

	ret, _ := suite.DBConn.GetTodo(todo.ID)
	assert.Equal(suite.T(), *todo, *ret)

	_ = suite.DBConn.DeleteTodo(todo.ID)
	ret, err = suite.DBConn.GetTodo(todo.ID)
	assert.NotNil(suite.T(), err)
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
