package db

import (
	"strconv"
	"strings"
)

// GENERAL DATA STRUCTS
type option struct {
	Opt string `sql:"opt" json:"opt"`
	Qty string `json:"qty"`
}

// Select struct for a simple select form
type Select struct {
	Title string   `json:"title"`
	Opts  []option `sql:"opts" json:"opts"`
}

func (sel *Select) parseSimple(field string, model interface{}) error {
	err := DBCon.Model(model).
		ColumnExpr(field + " as opt, count(*) as qty").
		Group(field).Select(&sel.Opts)
	if err != nil {
		return nil
	}
	return err
}

// SelectOption struct for a select with custom option
type SelectOption struct {
	Title  string   `json:"title"`
	Opts   []option `sql:"opts" json:"opts"`
	Others []string `sql:"others" json:"others"`
}

func (selopt *SelectOption) parse(field string, model interface{}) error {
	err := DBCon.Model(model).
		ColumnExpr(field + " as opt, count(*) as qty").
		Where(field + " ~ '^[0-9]+$'").
		Group(field).
		Select(&selopt.Opts)
	if err != nil {
		return err
	}
	err = DBCon.Model(model).
		Column(field).
		Where(field + " !~ '^[0-9]+$'").
		Select(&selopt.Others)
	if err != nil {
		return err
	}
	return nil
}

// Multiple handles response struct to multiple options
type Multiple struct {
	Title  string         `json:"title"`
	Opts   map[string]int `json:"opts"`
	Others []string       `sql:"data" json:"others"`
}

func (selopt *SelectOption) parseMul(field string, model interface{}) error {
	opts := make(map[string]int)
	var queryResults = []string{}
	err := DBCon.Model(model).
		Column(field).Select(&queryResults)
	if err != nil {
		return err
	}
	for _, dat := range queryResults {
		options := strings.Split(dat, ";")
		// we make the process when lenght is 2, because the answers is empty otherwise
		if len(options) == 2 {
			numbers := strings.Split(options[0], ",")
			// number holds the option number, send it to opts map
			for _, number := range numbers {
				if len(number) <= 1 {
					if num, ok := opts[number]; ok {
						opts[number] = num + 1
					} else {
						opts[number] = 1
					}
				}
			}
			freeOptions := strings.Split(options[1], ",")
			for _, freeOption := range freeOptions {
				if freeOption != "" {
					selopt.Others = append(selopt.Others, freeOption)
				}
			}
		}
	}
	for key, val := range opts {
		selopt.Opts = append(selopt.Opts, option{key, strconv.Itoa(val)})
	}
	return nil
}
