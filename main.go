package main

import (
	"net/http"

	"github.com/komangkrisnanda/golang-crud-mvc/controllers/patientcontroller"
)

func main() {
	http.HandleFunc("/", patientcontroller.Index)
	http.HandleFunc("/patient", patientcontroller.Index)
	http.HandleFunc("/patient/index", patientcontroller.Index)

	http.HandleFunc("/patient/add", patientcontroller.Add)
	http.HandleFunc("/patient/edit", patientcontroller.Edit)
	http.HandleFunc("/patient/delete", patientcontroller.Delete)
	
	http.ListenAndServe(":3000", nil)
}