package main

import (
	hooks_postgres "hook_pipe/internal/db/postgres/hooks"

	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hook_pipe v0.0.1")

	// Conecta a la base de datos Postgres
	db, err := gorm.Open(sqlite.Open("hooks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&hooks_postgres.HookModel{})
}
