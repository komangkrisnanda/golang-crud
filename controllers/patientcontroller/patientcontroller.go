package patientcontroller

import (
	"net/http"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request){
	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}


func Add(response http.ResponseWriter, request *http.Request){
	temp, err := template.ParseFiles("views/patient/add.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}


func Edit(response http.ResponseWriter, request *http.Request){
	
}


func Delete(response http.ResponseWriter, request *http.Request){
	
}
