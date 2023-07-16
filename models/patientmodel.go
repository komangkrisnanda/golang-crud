package models

import (
	"database/sql"
	"fmt"

	"github.com/komangkrisnanda/golang-crud-mvc/config"
	"github.com/komangkrisnanda/golang-crud-mvc/entities"
)

type PatientModel struct {
	conn *sql.DB
}

func NewPatientModel() *PatientModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PatientModel{
		conn: conn,
	}
}

func (p *PatientModel) FindAll() ([]entities.Patient, error){
	rows, err := p.conn.Query("SELECT * FROM patients")
	if err != nil {
		return []entities.Patient{}, err
	}
	defer rows.Close()

	var patients []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		rows.Scan(&patient.Id, &patient.Fullname, &patient.IdentityNumber, &patient.Gender, &patient.Pob, &patient.Dob, &patient.Address, &patient.Phone)

		patients = append(patients, patient)
	}

	return patients, nil
}

func (p *PatientModel) Create(patient entities.Patient) bool{
	result, err := p.conn.Exec("INSERT INTO patients (fullname, identity_number, gender, pob, dob, address, phone) VALUES (?, ?, ?, ?, ?, ?, ?) ", patient.Fullname, patient.IdentityNumber, patient.Gender,  patient.Pob, patient.Dob, patient.Address, patient.Phone)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}