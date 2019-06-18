package db

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Student represents the student table
type Student struct {
	refPointer int      `sql:"-"`
	tableName  struct{} `sql:"student"`
	Ci         string   `sql:"ci,pk" json:"ci"`
	ApPat      string   `sql:"appat" json:"appat"`
	ApMat      string   `sql:"apmat" json:"apmat"`
	Nombre     string   `sql:"nombre" json:"nombre"`
	PerIngreso string   `sql:"peringreso" json:"perIngreso"`
	FecNac     string   `sql:"fecnac" json:"fecNac"`
	Telf       string   `sql:"telf" json:"telf"`
	Cel        string   `sql:"cel" json:"cel"`
	Email      string   `sql:"email" json:"email"`
	Email2     string   `sql:"email2" json:"email2"`
	Genero     string   `sql:"genero" json:"genero"`
	Cole       string   `sql:"cole" json:"cole"`
	TipoCole   string   `sql:"tipocole" json:"tipoCole"`
	ModIngreso string   `sql:"modingreso" json:"modIngreso"`
}

// CreateStudentTable creates a student table
func CreateStudentTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Student{}, opts)
	if err != nil {
		log.Printf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table Student created.")
	return nil
}

// GetByCi get user by email
func (u *Student) GetByCi(ci string) error {
	err := DBCon.Model(u).Where("ci = ?", ci).Select()
	if err != nil {
		return err
	}
	return nil
}
