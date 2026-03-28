package entities

type Recruiter struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	License_no int    `json:"license_no"`
	Location   string `json:"location"`
	Phone_no   int    `json:"phone_no"`
}
