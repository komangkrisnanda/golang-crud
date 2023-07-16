package patientcontroller

import (
	"net/http"
	"text/template"

	"github.com/komangkrisnanda/golang-crud-mvc/entities"
	"github.com/komangkrisnanda/golang-crud-mvc/models"
)

var patientModel = models.NewPatientModel()

func Index(response http.ResponseWriter, request *http.Request){

	patients, _ := patientModel.FindAll()

	data := map[string]interface{}{
		"patients": patients,
	}

	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}


func Add(response http.ResponseWriter, request *http.Request){

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	}else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient
		patient.Fullname = request.Form.Get("fullname")
		patient.IdentityNumber = request.Form.Get("identityNumber")
		patient.Gender = request.Form.Get("gender")
		patient.Pob = request.Form.Get("pob")
		patient.Dob = request.Form.Get("dob")
		patient.Address = request.Form.Get("address")
		patient.Phone = request.Form.Get("phoneNumber")
		
		patientModel.Create(patient)

		data := map[string]interface{}{
			"message" : "Data successfully stored.",
		}
		
		temp, _ := template.ParseFiles("views/patient/add.html")
		
		temp.Execute(response, data)

	}
	

}


func Edit(response http.ResponseWriter, request *http.Request){
	
}


func Delete(response http.ResponseWriter, request *http.Request){
	
}
