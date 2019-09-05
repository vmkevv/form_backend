package db

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
