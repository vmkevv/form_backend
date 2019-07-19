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
	Type       string    `sql:"type" json:"type"`
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

// GetByID get user by id
func (u *User) GetByID() error {
	err := DBCon.Select(u)
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
	err := db.Select(u)
	if err != nil {
		return err
	}
	u.IsActive = !u.IsActive
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

// UpdatePassword updates the password
func (u *User) UpdatePassword(newPass string) error {
	u.Password = newPass
	if err := DBCon.Update(u); err != nil {
		return err
	}
	return nil
}

// GetAll return a list of books
func GetAll() ([]User, error) {
	var users []User
	err := DBCon.Model(&users).Order("id").Where("type = ?", "user").Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetFormsList gets all forms of an user
func (u *User) GetFormsList() (interface{}, error) {
	type form struct {
		ID        int       `json:"id"`
		Nro       string    `json:"nro"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	type respStruct struct {
		EstForms []form `json:"form-est"`
		DocForms []form `json:"form-doc"`
		ProForms []form `json:"form-pro"`
		PreForms []form `json:"form-pre"`
	}
	resp := respStruct{[]form{}, []form{}, []form{}, []form{}}

	var estForms []FormEst
	errEst := DBCon.Model(&estForms).Where("user_id = ?", u.ID).Column("id", "nro", "updated_at").Select()
	if errEst != nil {
		return nil, errEst
	}
	for _, estForm := range estForms {
		resp.EstForms = append(resp.EstForms, form{estForm.ID, estForm.Nro, estForm.UpdatedAt})
	}

	var docForms []FormDoc
	errDoc := DBCon.Model(&docForms).Where("user_id = ?", u.ID).Column("id", "nro", "updated_at").Select()
	if errDoc != nil {
		return nil, errDoc
	}
	for _, docForm := range docForms {
		resp.DocForms = append(resp.DocForms, form{docForm.ID, docForm.Nro, docForm.UpdatedAt})
	}

	var proForms []FormPro
	errPro := DBCon.Model(&proForms).Where("user_id = ?", u.ID).Column("id", "nro", "updated_at").Select()
	if errPro != nil {
		return nil, errPro
	}
	for _, proForm := range proForms {
		resp.ProForms = append(resp.ProForms, form{proForm.ID, proForm.Nro, proForm.UpdatedAt})
	}

	var preForms []FormPre
	errPre := DBCon.Model(&preForms).Where("user_id = ?", u.ID).Column("id", "nro", "updated_at").Select()
	if errPre != nil {
		return nil, errPre
	}
	for _, preForm := range preForms {
		resp.PreForms = append(resp.PreForms, form{preForm.ID, preForm.Nro, preForm.UpdatedAt})
	}

	return resp, nil
}
