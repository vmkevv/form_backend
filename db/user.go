package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
)

// User represents user table in database
type User struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"user"`
	ID         int       `sql:"id,pk" json:"id"`
	Name       string    `sql:"nombre" json:"nombre"`
	LastName1  string    `sql:"appat" json:"appat"`
	LastName2  string    `sql:"apmat" json:"apmat"`
	Ci         string    `sql:"ci" json:"ci"`
	Phone      int       `sql:"cel" json:"cel"`
	Email      string    `sql:"correo" json:"correo"`
	Password   string    `sql:"password" json:"-"`
	HomeDir    string    `sql:"dir" json:"dir"`
	CreatedAt  time.Time `sql:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at"`
	IsActive   bool      `sql:"activo" json:"activo"`
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

// GetByEmail get user by email
func (u *User) GetByEmail(correo string) error {
	err := DBCon.Model(u).Where("correo = ?", correo).Select()
	if err != nil {
		return err
	}
	return nil
}

// CheckEmail check if given email exists
func (u *User) CheckEmail(db *pg.DB) bool {
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		fmt.Println(err.Error())
	}
	var existEmail = false
	for _, user := range users {
		if u.Email == user.Email {
			existEmail = true
		}
	}
	return existEmail
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

// ChangeActive changes active values
func (u *User) ChangeActive(db *pg.DB) error {
	updated := u.IsActive
	err := db.Select(u)
	if err != nil {
		return err
	}
	u.IsActive = updated
	err = db.Update(u)
	if err != nil {
		return err
	}
	return nil
}

// ResetPassword reset password to CI
func (u *User) ResetPassword(db *pg.DB) error {
	err := db.Select(u)
	if err != nil {
		return err
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Ci), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	err = db.Update(u)
	if err != nil {
		return err
	}
	return nil
}

// GetAll return a list of books
func GetAll() ([]User, error) {
	var users []User
	err := DBCon.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
