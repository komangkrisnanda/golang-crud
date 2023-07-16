package patientcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/komangkrisnanda/golang-crud-mvc/entities"
	"github.com/komangkrisnanda/golang-crud-mvc/libraries"
	"github.com/komangkrisnanda/golang-crud-mvc/models"
)

var validation = libraries.NewValidation()
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

		var data = make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		}else{
			data["message"] = "Data successfully stored."
			patientModel.Create(patient)
		}
		
		temp, _ := template.ParseFiles("views/patient/add.html")
		
		temp.Execute(response, data)

	}
	

}


func Edit(response http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var patient entities.Patient
		patientModel.Find(id, &patient)

		data := map[string]interface{}{
			"patient": patient,
		}

		temp, err := template.ParseFiles("views/patient/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	}else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient

		id, _:= strconv.ParseInt(request.Form.Get("id"), 10, 64)

		patient.Id = id
		patient.Fullname = request.Form.Get("fullname")
		patient.IdentityNumber = request.Form.Get("identityNumber")
		patient.Gender = request.Form.Get("gender")
		patient.Pob = request.Form.Get("pob")
		patient.Dob = request.Form.Get("dob")
		patient.Address = request.Form.Get("address")
		patient.Phone = request.Form.Get("phoneNumber")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		}else{
			data["message"] = "Data successfully updated."
			patientModel.Update(patient)
		}
		
		// temp, _ := template.ParseFiles("views/patient/index.html")
		
		// temp.Execute(response, data)

		http.Redirect(response, request, "/patient", http.StatusSeeOther)	

	}
}


func Delete(response http.ResponseWriter, request *http.Request){
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	patientModel.Delete(id)

	http.Redirect(response, request, "/patient", http.StatusSeeOther)
}
