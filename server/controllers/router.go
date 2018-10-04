package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kansi/go_wasm_todo/server/models"
	"net/http"
)

// Router for the REST api
func Router(db *models.Db) *gin.Engine {
	r := gin.Default()

	// CLIENT
	// The mime type for wasm file needs to specifically set
	// so that the browser can process it correctly
	{
		r.StaticFile("/", "../asset/wasm_exec.html")
		r.StaticFS("/asset", http.Dir("../asset"))
		r.StaticFS("/sui", http.Dir("../sui"))

		r.GET("/client.wasm", func(c *gin.Context) {
			c.Header("Content-Type", "application/wasm")
			c.Writer.WriteHeaderNow()
			c.File("../asset/client.wasm")
		})
	}

	tc := NewTodoController(db)
	// Server REST routes
	{
		r.GET("/todo/:id", tc.GetTodo)
		r.PUT("/todo/:id", tc.StoreTodo)
		r.DELETE("/todo/:id", tc.DeleteTodo)
	}
	return r
}
