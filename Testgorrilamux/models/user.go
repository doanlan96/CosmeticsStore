package models

type User struct {
	Id          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    []byte `json:"password"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Avatar      string `json:"avatar"`
	RoleId      uint   `json:"role_id"`
}

type Customer struct {
	Id          uint   			`json:"id"`
	FullName    string 			`json:"full_name"`
	Email       string 			`json:"email"`
	Phone       string 			`json:"phone"`
	Address     string 			`json:"address"`
	DateOfBirth string 			`json:"date_of_birth"`
	Gender      string 			`json:"gender"`
	Avatar      string 			`json:"avatar"`
	RoleId      uint   			`json:"role_id"`
	Order 		[]OrderItem 	`json:"order"`
}
