package main

import (
	"errors"
	"fmt"
	"sort"
)

// Database holds multiple named todo tables and the currently active table name
type Database struct {
	Tables  map[string]*Todos `json:"Tables"`
	Current string            `json:"Current"`
}

func NewDatabase() *Database {
	d := &Database{}
	d.ensureDefault()
	return d
}

func (d *Database) ensureDefault() {
	if d.Tables == nil {
		d.Tables = make(map[string]*Todos)
	}
	if d.Current == "" {
		d.Current = "default"
	}
	if _, ok := d.Tables[d.Current]; !ok {
		d.Tables[d.Current] = &Todos{}
	}
}

// GetCurrent returns a pointer to the currently active Todos slice
func (d *Database) GetCurrent() *Todos {
	d.ensureDefault()
	return d.Tables[d.Current]
}

// NewTable creates a new table and switches to it. Returns an error if the table already exists.
func (d *Database) NewTable(name string) error {
	if name == "" {
		return errors.New("table name cannot be empty")
	}
	d.ensureDefault()
	if _, ok := d.Tables[name]; ok {
		return errors.New("table already exists")
	}
	d.Tables[name] = &Todos{}
	d.Current = name
	return nil
}

// SwitchTable switches to an existing table
func (d *Database) SwitchTable(name string) error {
	if name == "" {
		return errors.New("table name cannot be empty")
	}
	d.ensureDefault()
	if _, ok := d.Tables[name]; !ok {
		return errors.New("table not found")
	}
	d.Current = name
	return nil
}

// AllTableNames returns the list of table names sorted
func (d *Database) AllTableNames() []string {
	d.ensureDefault()
	names := make([]string, 0, len(d.Tables))
	for k := range d.Tables {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

// PrintTables prints the list of tables and marks the current one
func (d *Database) PrintTables() {
	names := d.AllTableNames()
	if len(names) == 0 {
		fmt.Println("No tables found")
		return
	}
	for _, n := range names {
		marker := " "
		if n == d.Current {
			marker = "*"
		}
		fmt.Printf("%s %s\n", marker, n)
	}
}

// MigrateFromOld loads a Todos value into the default table (used when migrating old file format)
func (d *Database) MigrateFromOld(t Todos) {
	d.ensureDefault()
	if t == nil {
		return
	}
	// if default table already has entries, append to them
	*d.Tables["default"] = append(*d.Tables["default"], t...)
}
