package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// User represents user table in database
type User struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"user"`
	ID         int       `sql:"id,pk"`
	Name       string    `sql:"nombre"`
	LastName1  string    `sql:"appat"`
	LastName2  string    `sql:"apmat"`
	Ci         string    `sql:"ci"`
	Phone      int       `sql:"cel"`
	Email      string    `sql:"correo"`
	Password   string    `sql:"password"`
	HomeDir    string    `sql:"dir"`
	CreatedAt  time.Time `sql:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at"`
	isActive   bool      `sql:"activo"`
}

// Save saves user struct
func (u *User) Save(db *pg.DB) error {
	err := db.Insert(u)
	if err != nil {
		log.Printf("Error insert data, %v", err)
		return err
	}
	return nil
}

// CreateUserTable creates a user table
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&User{}, opts)
	if err != nil {
		log.Printf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table created.")
	return nil
}
