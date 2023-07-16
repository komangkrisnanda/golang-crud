package entities

type Patient struct {
	Id 				int64
	Fullname 		string `validate:"required" label:"Fullname"`
	IdentityNumber 	string `validate:"required" label:"Identity Number"`
	Gender 			string `validate:"required" label:"Gender"`
	Pob 			string `validate:"required" label:"Place of Birth"`
	Dob 			string `validate:"required" label:"Date of Birth"`
	Address 		string `validate:"required" label:"Address"`
	Phone 			string `validate:"required" label:"Phone Number"`
}