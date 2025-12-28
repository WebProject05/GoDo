package main

import "fmt"

func main() {
	db := NewDatabase()
	storage := NewStorage[Database]("todos.json")
	if err := storage.Load(db); err != nil {
		// try migrating old format (plain Todos) into the default table
		var old Todos
		oldStorage := NewStorage[Todos]("todos.json")
		if err2 := oldStorage.Load(&old); err2 == nil {
			db.MigrateFromOld(old)
			fmt.Println("Migrated existing todos into table 'default'")
		} else {
			// file doesn't exist or cannot be read, keep fresh DB
		}
	}

	cmdFlags := newCmdFlags()
	cmdFlags.Execute(db)

	if err := storage.Save(*db); err != nil {
		fmt.Println("Error saving todos:", err)
	}
}
