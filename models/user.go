package models

//User structure with 5 sample fields
type User struct {
	UUID      string `json:id`
	MobileID  string `json:"mobile_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
