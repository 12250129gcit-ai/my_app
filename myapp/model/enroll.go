package model

import (
	"myapp/dataStore/postgres"
)

type Enroll struct {
	StdId        int64  `json:"stdid"`
	CourseID     string `json:"cid"`
	DateEnrolled string `json:"date_enrolled"`
}

// Insert a new enrollment
const queryInsertEnroll = `INSERT INTO enroll(std_id, course_id, date_enrolled) VALUES($1, $2, $3) RETURNING std_id`

func (e *Enroll) EnrollStud() error {
	return postgres.Db.QueryRow(queryInsertEnroll, e.StdId, e.CourseID, e.DateEnrolled).Scan(&e.StdId)
}

// Get a single enrollment
const queryGetEnroll = `SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 AND course_id=$2`

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.DateEnrolled)
}

// Delete an enrollment
const queryDeleteEnroll = `DELETE FROM enroll WHERE std_id=$1 AND course_id=$2 RETURNING std_id`

func (e *Enroll) Delete() error {
	return postgres.Db.QueryRow(queryDeleteEnroll, e.StdId, e.CourseID).Scan(&e.StdId)
}

// Get all enrollments with student and course details
func GetAllEnrollments() ([]map[string]interface{}, error) {
	query := `
        SELECT s.stdid, s.firstname, s.lastname, c.cid, c.coursename, e.date_enrolled 
        FROM enroll e
        JOIN student s ON e.std_id = s.stdid
        JOIN course c ON e.course_id = c.cid
        ORDER BY e.date_enrolled DESC
    `

	rows, err := postgres.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []map[string]interface{}
	for rows.Next() {
		var stdid int64
		var firstname, lastname, cid, coursename, date_enrolled string

		err := rows.Scan(&stdid, &firstname, &lastname, &cid, &coursename, &date_enrolled)
		if err != nil {
			return nil, err
		}

		enrollment := map[string]interface{}{
			"stdid":         stdid,
			"student_name":  firstname + " " + lastname,
			"cid":           cid,
			"course_name":   coursename,
			"date_enrolled": date_enrolled,
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
