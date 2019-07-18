package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormDoc doc form structure
type FormDoc struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_doc"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Doc1       string    `sql:"doc1" json:"doc1"`
	Doc2       string    `sql:"doc2" json:"doc2"`
	Doc3       string    `sql:"doc3" json:"doc3"`
	Doc4       string    `sql:"doc4" json:"doc4"`
	Doc5       string    `sql:"doc5" json:"doc5"`
	Doc6       string    `sql:"doc6" json:"doc6"`
	Doc7       string    `sql:"doc7" json:"doc7"`
	Doc8       string    `sql:"doc8" json:"doc8"`
	Doc9       string    `sql:"doc9" json:"doc9"`
	Doc10      string    `sql:"doc10" json:"doc10"`
	Doc11      string    `sql:"doc11" json:"doc11"`
	Doc12      string    `sql:"doc12" json:"doc12"`
	Doc13      string    `sql:"doc13" json:"doc13"`
	Doc14a     string    `sql:"doc14a" json:"doc14a"`
	Doc14b     string    `sql:"doc14b" json:"doc14b"`
	Doc14c     string    `sql:"doc14c" json:"doc14c"`
	Doc14d     string    `sql:"doc14d" json:"doc14d"`
	Doc15      string    `sql:"doc15" json:"doc15"`
	Doc16      string    `sql:"doc16" json:"doc16"`
	Doc17      string    `sql:"doc17" json:"doc17"`
	Doc18      string    `sql:"doc18" json:"doc18"`
	Doc19      string    `sql:"doc19" json:"doc19"`
	Doc20      string    `sql:"doc20" json:"doc20"`
	Doc21      string    `sql:"doc21" json:"doc21"`
	Doc22      string    `sql:"doc22" json:"doc22"`
	Doc23      string    `sql:"doc23" json:"doc23"`
	Doc24      string    `sql:"doc24" json:"doc24"`
	Doc24a     string    `sql:"doc24a" json:"doc24a"`
	Doc25      string    `sql:"doc25" json:"doc25"`
	Doc25a     string    `sql:"doc25a" json:"doc25a"`
	Doc26      string    `sql:"doc26" json:"doc26"`
	Doc26a     string    `sql:"doc26a" json:"doc26a"`
	Doc27      string    `sql:"doc27" json:"doc27"`
	Doc27a     string    `sql:"doc27a" json:"doc27a"`
	Doc28      string    `sql:"doc28" json:"doc28"`
	Doc29      string    `sql:"doc29" json:"doc29"`
	Doc30      string    `sql:"doc30" json:"doc30"`
	Doc31      string    `sql:"doc31" json:"doc31"`
	Doc32      string    `sql:"doc32" json:"doc32"`
	Doc33      string    `sql:"doc33" json:"doc33"`
	Doc34      string    `sql:"doc34" json:"doc34"`
	Doc35      string    `sql:"doc35" json:"doc35"`
	Doc36      string    `sql:"doc36" json:"doc36"`
	Doc37      string    `sql:"doc37" json:"doc37"`
	Doc38      string    `sql:"doc38" json:"doc38"`
	Doc39      string    `sql:"doc39" json:"doc39"`
	Doc40      string    `sql:"doc40" json:"doc40"`
	Doc41      string    `sql:"doc41" json:"doc41"`
	Doc42      string    `sql:"doc42" json:"doc42"`
	Doc43      string    `sql:"doc43" json:"doc43"`
	Doc44      string    `sql:"doc44" json:"doc44"`
	Doc45      string    `sql:"doc45" json:"doc45"`
	Doc46      string    `sql:"doc46" json:"doc46"`
	Doc47      string    `sql:"doc47" json:"doc47"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// CreateFormDocTable creates doc table
func CreateFormDocTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	if err := db.CreateTable(&FormDoc{}, opts); err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormDoc created")
	return nil
}

// Save saves the doc form
func (fdoc *FormDoc) Save() error {
	err := DBCon.Insert(fdoc)
	if err != nil {
		return err
	}
	return nil
}

// GetByNro get doc form by nro
func (fdoc *FormDoc) GetByNro(nro string) error {
	err := DBCon.Model(fdoc).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}
