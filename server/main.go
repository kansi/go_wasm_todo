package main

import (
	"github.com/kansi/go_wasm_todo/server/controllers"
	"github.com/kansi/go_wasm_todo/server/models"
)

func main() {
	// Setup database connection
	db := models.NewDbConnection("/tmp/go_wasm_todo_prod.db")

	// Start server
	r := controllers.Router(db)
	r.Run(":8000")
}
