package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormIns institutions form
type FormIns struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_ins"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Ins1       string    `sql:"ins1" json:"ins1"`
	Ins2       string    `sql:"ins2" json:"ins2"`
	Ins3       string    `sql:"ins3" json:"ins3"`
	Ins4       string    `sql:"ins4" json:"ins4"`
	Ins5       string    `sql:"ins5" json:"ins5"`
	Ins6       string    `sql:"ins6" json:"ins6"`
	Ins7       string    `sql:"ins7" json:"ins7"`
	Ins8       string    `sql:"ins8" json:"ins8"`
	Ins9       string    `sql:"ins9" json:"ins9"`
	Ins10      string    `sql:"ins10" json:"ins10"`
	Ins11      string    `sql:"ins11" json:"ins11"`
	Ins12      string    `sql:"ins12" json:"ins12"`
	Ins13a     string    `sql:"ins13a" json:"ins13a"`
	Ins13b     string    `sql:"ins13b" json:"ins13b"`
	Ins13c     string    `sql:"ins13c" json:"ins13c"`
	Ins13d     string    `sql:"ins13d" json:"ins13d"`
	Ins13e     string    `sql:"ins13e" json:"ins13e"`
	Ins13f     string    `sql:"ins13f" json:"ins13f"`
	Ins13g     string    `sql:"ins13g" json:"ins13g"`
	Ins13h     string    `sql:"ins13h" json:"ins13h"`
	Ins13i     string    `sql:"ins13i" json:"ins13i"`
	Ins13j     string    `sql:"ins13j" json:"ins13j"`
	Ins13k     string    `sql:"ins13k" json:"ins13k"`
	Ins14      string    `sql:"ins14" json:"ins14"`
	Ins15      string    `sql:"ins15" json:"ins15"`
	Ins16      string    `sql:"ins16" json:"ins16"`
	Ins17      string    `sql:"ins17" json:"ins17"`
	Ins18      string    `sql:"ins18" json:"ins18"`
	Ins19      string    `sql:"ins19" json:"ins19"`
	Ins20      string    `sql:"ins20" json:"ins20"`
	Ins21      string    `sql:"ins21" json:"ins21"`
	Ins22      string    `sql:"ins22" json:"ins22"`
	Ins23      string    `sql:"ins23" json:"ins23"`
	Ins24      string    `sql:"ins24" json:"ins24"`
	Ins25      string    `sql:"ins25" json:"ins25"`
	Ins26      string    `sql:"ins26" json:"ins26"`
	Ins27      string    `sql:"ins27" json:"ins27"`
	Ins28      string    `sql:"ins28" json:"ins28"`
	Ins29      string    `sql:"ins29" json:"ins29"`
	Ins30      string    `sql:"ins30" json:"ins30"`
	Ins31      string    `sql:"ins31" json:"ins31"`
	Ins32      string    `sql:"ins32" json:"ins32"`
	Ins33      string    `sql:"ins33" json:"ins33"`
	Ins34      string    `sql:"ins34" json:"ins34"`
	Ins35      string    `sql:"ins35" json:"ins35"`
	Ins36      string    `sql:"ins36" json:"ins36"`
	Ins37      string    `sql:"ins37" json:"ins37"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// Save saves the institution form
func (fins *FormIns) Save() error {
	if err := DBCon.Insert(fins); err != nil {
		return err
	}
	return nil
}

// CreateFormInsTable creates the table
func CreateFormInsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormIns{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormIns created")
	return nil
}

// Update updates the form
func (fins *FormIns) Update() error {
	err := DBCon.Update(fins)
	if err != nil {
		return err
	}
	return nil
}

// GetByNro get form by nro
func (fins *FormIns) GetByNro(nro string) error {
	err := DBCon.Model(fins).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}
