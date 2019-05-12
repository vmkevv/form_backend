package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Admin represents user table in database
type Admin struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"admin"`
	ID         int       `sql:"id,pk"`
	Name       string    `sql:"nombre"`
	Email      string    `sql:"correo"`
	Password   string    `sql:"password"`
	CreatedAt  time.Time `sql:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at"`
}

// Save saves user struct
func (a *Admin) Save(db *pg.DB) error {
	err := db.Insert(a)
	if err != nil {
		log.Printf("Error insert data, %v", err)
		return err
	}
	return nil
}

// GetByEmail get by email
func (a *Admin) GetByEmail(correo string) error {
	err := DBCon.Model(a).Where("correo = ?", correo).Select()
	if err != nil {
		return err
	}
	return nil
}

// CreateAdminTable creates a user table
func CreateAdminTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Admin{}, opts)
	if err != nil {
		log.Printf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table created.")
	return nil
}
