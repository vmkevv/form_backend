package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormPro profesionals form
type FormPro struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_pro"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Pro1       string    `sql:"pro1" json:"pro1"`
	Pro2       string    `sql:"pro2" json:"pro2"`
	Pro3       string    `sql:"pro3" json:"pro3"`
	Pro4       string    `sql:"pro4" json:"pro4"`
	Pro5       string    `sql:"pro5" json:"pro5"`
	Pro6       string    `sql:"pro6" json:"pro6"`
	Pro7       string    `sql:"pro7" json:"pro7"`
	Pro8       string    `sql:"pro8" json:"pro8"`
	Pro9       string    `sql:"pro9" json:"pro9"`
	Pro10      string    `sql:"pro10" json:"pro10"`
	Pro11      string    `sql:"pro11" json:"pro11"`
	Pro12      string    `sql:"pro12" json:"pro12"`
	Pro13      string    `sql:"pro13" json:"pro13"`
	Pro14      string    `sql:"pro14" json:"pro14"`
	Pro15      string    `sql:"pro15" json:"pro15"`
	Pro16      string    `sql:"pro16" json:"pro16"`
	Pro17      string    `sql:"pro17" json:"pro17"`
	Pro18      string    `sql:"pro18" json:"pro18"`
	Pro19      string    `sql:"pro19" json:"pro19"`
	Pro20      string    `sql:"pro20" json:"pro20"`
	Pro21      string    `sql:"pro21" json:"pro21"`
	Pro22      string    `sql:"pro22" json:"pro22"`
	Pro23      string    `sql:"pro23" json:"pro23"`
	Pro24      string    `sql:"pro24" json:"pro24"`
	Pro25      string    `sql:"pro25" json:"pro25"`
	Pro26      string    `sql:"pro26" json:"pro26"`
	Pro27      string    `sql:"pro27" json:"pro27"`
	Pro28      string    `sql:"pro28" json:"pro28"`
	Pro29      string    `sql:"pro29" json:"pro29"`
	Pro30      string    `sql:"pro30" json:"pro30"`
	Pro31      string    `sql:"pro31" json:"pro31"`
	Pro32      string    `sql:"pro32" json:"pro32"`
	Pro33      string    `sql:"pro33" json:"pro33"`
	Pro34      string    `sql:"pro34" json:"pro34"`
	Pro35      string    `sql:"pro35" json:"pro35"`
	Pro36      string    `sql:"pro36" json:"pro36"`
	Pro37      string    `sql:"pro37" json:"pro37"`
	Pro38      string    `sql:"pro38" json:"pro38"`
	Pro39      string    `sql:"pro39" json:"pro39"`
	Pro40      string    `sql:"pro40" json:"pro40"`
	Pro41      string    `sql:"pro41" json:"pro41"`
	Pro42      string    `sql:"pro42" json:"pro42"`
	Pro43      string    `sql:"pro43" json:"pro43"`
	Pro44      string    `sql:"pro44" json:"pro44"`
	Pro45      string    `sql:"pro45" json:"pro45"`
	Pro46      string    `sql:"pro46" json:"pro46"`
	Pro47      string    `sql:"pro47" json:"pro47"`
	Pro48      string    `sql:"pro48" json:"pro48"`
	Pro49      string    `sql:"pro49" json:"pro49"`
	Pro50      string    `sql:"pro50" json:"pro50"`
	Pro51      string    `sql:"pro51" json:"pro51"`
	Pro52      string    `sql:"pro52" json:"pro52"`
	Pro53      string    `sql:"pro53" json:"pro53"`
	Pro54      string    `sql:"pro54" json:"pro54"`
	Pro55      string    `sql:"pro55" json:"pro55"`
	Pro56      string    `sql:"pro56" json:"pro56"`
	Pro57      string    `sql:"pro57" json:"pro57"`
	Pro58      string    `sql:"pro58" json:"pro58"`
	Pro59      string    `sql:"pro59" json:"pro59"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// Save saves the profesional form
func (fpro *FormPro) Save() error {
	err := DBCon.Insert(fpro)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

// CreateFormProTable creates the table
func CreateFormProTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormPro{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormPro created")
	return nil
}

// GetByNro get form by nro
func (fpro *FormPro) GetByNro(nro string) error {
	err := DBCon.Model(fpro).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the profesiona form
func (fpro *FormPro) Update() error {
	var formProAux FormPro
	formProAux.ID = fpro.ID
	if err := formProAux.GetByID(); err != nil {
		return err
	}
	fpro.UserID = formProAux.UserID
	fpro.CreatedAt = formProAux.CreatedAt
	fpro.UpdatedAt = time.Now()
	if err := DBCon.Update(fpro); err != nil {
		return err
	}
	return nil
}

// Delete deletes the profesional form
func (fpro *FormPro) Delete() error {
	if err := DBCon.Delete(fpro); err != nil {
		return err
	}
	return nil
}

// GetByID get form by id
func (fpro *FormPro) GetByID() error {
	if err := DBCon.Select(fpro); err != nil {
		return err
	}
	return nil
}
