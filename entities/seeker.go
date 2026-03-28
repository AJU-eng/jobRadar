package entities

type Seeker struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	Qualification string `json:"qualification"`
	Adhar_no      int    `json:"adhar_no"`
	Phone_no      int    `json:"phone_no"`
	Location      string `json:"location"`
}
