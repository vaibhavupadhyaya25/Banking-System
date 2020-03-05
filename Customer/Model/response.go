package Model

//CustomerRes ..
type CustomerRes struct {
	tableName struct{} `pg:"customer"`
	Fname     string   `json:"fname,omitempty" validate:"required"`
	Lname     string   `json:"lname,omitempty" validate:"required"`
	Address   string   `json:"address,omitempty" validate:"required"`
	Phoneno   string   `json:"phoneno,omitempty" validate:"required,min=10,max=12"`
	Emailid   string   `json:"emailid,omitempty" validate:"required,email"`
}
