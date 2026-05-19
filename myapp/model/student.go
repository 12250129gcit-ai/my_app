package model

import (
	"myapp/dataStore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"` // Changed from "firstname" to "fname"
	LastName  string `json:"lname"` // Changed from "lastname" to "lname"
	Email     string `json:"email"`
}

// Create student
const queryInsert = `INSERT INTO student (stdid, firstname, lastname, email)
                     VALUES ($1, $2, $3, $4)`

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsert, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

// Read one student by ID
const queryGet = `SELECT stdid, firstname, lastname, email FROM student WHERE stdid = $1`

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGet, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

// Update student
const queryUpdate = `UPDATE student SET stdid = $1, firstname = $2, lastname = $3, email = $4
                     WHERE stdid = $5 RETURNING stdid`

func (s *Student) Update(oldId int64) error {
	err := postgres.Db.QueryRow(queryUpdate, s.StdId, s.FirstName, s.LastName, s.Email, oldId).Scan(&s.StdId)
	return err
}

// Delete student
const queryDelete = `DELETE FROM student WHERE stdid = $1 RETURNING stdid`

func (s *Student) Delete() error {
	err := postgres.Db.QueryRow(queryDelete, s.StdId).Scan(&s.StdId)
	return err
}

// Get all students
func GetAllStudents() ([]Student, error) {
	rows, err := postgres.Db.Query("SELECT stdid, firstname, lastname, email FROM student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []Student{}
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
