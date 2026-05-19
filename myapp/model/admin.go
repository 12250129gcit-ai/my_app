package model

import (
	"myapp/dataStore/postgres"
)

type Admin struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

const queryInsertAdmin = "INSERT INTO admin(firstname, lastname, email, password) VALUES($1, $2, $3, $4) RETURNING email;"
const queryGetAdmin = "SELECT email, password FROM admin WHERE email=$1 and password=$2;"

func (adm *Admin) Create() error {
	return postgres.Db.QueryRow(queryInsertAdmin, adm.FirstName, adm.LastName, adm.Email, adm.Password).Scan(&adm.Email)
}

func (adm *Admin) Get() error {
	return postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password).Scan(&adm.Email, &adm.Password)
}
package model

import (
    "myapp/dataStore/postgres"
)

type Admin struct {
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}

// Create new admin (signup)
const queryInsertAdmin = `INSERT INTO admin(firstname, lastname, email, password) VALUES($1, $2, $3, $4) RETURNING email`

func (adm *Admin) Create() error {
    return postgres.Db.QueryRow(queryInsertAdmin, adm.FirstName, adm.LastName, adm.Email, adm.Password).Scan(&adm.Email)
}

// Get admin by email and password (login)
const queryGetAdmin = `SELECT email, password FROM admin WHERE email=$1 AND password=$2`

func (adm *Admin) Get() error {
    return postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password).Scan(&adm.Email, &adm.Password)
}