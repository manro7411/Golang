package repository

import (
	"database/sql"
	"fmt"
)

type userRepositoryDB struct {
	db *sql.DB
}

//outer folder UPPERCASE variable
func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(email string, password string, secret string) (*User, error) {
	fmt.Println("CreateUser")
	return nil, nil
}

func (r userRepositoryDB) CheckUser(email string) (*User, error) {
	fmt.Printf("CheckUser")
	return nil, nil
}
