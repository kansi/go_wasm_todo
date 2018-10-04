package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kansi/go_wasm_todo/server/models"
	"net/http"
)

// TodoController methods for storing, retrieving and updating todos
type TodoController struct {
	DB *models.Db
}

// NewTodoController Create a new instance of Todo controller
func NewTodoController(db *models.Db) *TodoController {
	return &TodoController{DB: db}
}

// GetTodo get todo by Id
func (tc *TodoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("fetching id ", id)
	todo, err := tc.DB.GetTodo(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Todo not found"})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// StoreTodo update the given todo
func (tc *TodoController) StoreTodo(c *gin.Context) {
	// id := c.Param("id")
}

// DeleteTodo with ID
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := tc.DB.DeleteTodo(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error while deleteing Todo"})
	} else {
		c.JSON(200, gin.H{"message": "Todo deleted successfully"})
	}
}
