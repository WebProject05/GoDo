package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
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

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
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
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid Command. GO to help for details of commands")
	}

}
