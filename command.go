package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add       string
	Del       int
	Edit      string
	Toggle    int
	List      bool
	NewTable  string
	Switch    string
	AllTables bool
}

func newCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new Todo Task Title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the Todo Task by Index and Title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo index to Toggle")
	flag.BoolVar(&cf.List, "list", false, "List all the todos in a tabular form")
	// keep `show` as an alias for backward compatibility
	flag.BoolVar(&cf.List, "show", false, "List all the todos in a tabular form (alias)")

	flag.StringVar(&cf.NewTable, "newtable", "", "Create a new todo table and switch to it")
	flag.StringVar(&cf.Switch, "switch", "", "Switch to an existing todo table")
	flag.BoolVar(&cf.AllTables, "alltables", false, "List all available todo tables")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(db *Database) {
	switch {
	case cf.AllTables:
		db.PrintTables()
	case cf.NewTable != "":
		if err := db.NewTable(cf.NewTable); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Switched to new table:", cf.NewTable)
	case cf.Switch != "":
		if err := db.SwitchTable(cf.Switch); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Switched to table:", cf.Switch)
	case cf.List:
		fmt.Println("Table:", db.Current)
		db.GetCurrent().print()
	case cf.Add != "":
		db.GetCurrent().add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		if err := db.GetCurrent().edit(index, parts[1]); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case cf.Toggle != -1:
		if err := db.GetCurrent().toggle(cf.Toggle); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case cf.Del != -1:
		if err := db.GetCurrent().delete(cf.Del); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid Command. GO to help for details of commands")
	}

}
