package model

// User example
type User struct {
	BaseModel
	PhotoUrls []string `json:"photoUrls"`
	Person    Person   `json:"person"`
}

// Person example
type Person struct {
	BaseModel
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// BaseModel example
type BaseModel struct {
	ID        string `json:"id"`
	CreatedAt int
	UpdatedAt string
}

// APIError example
type APIError struct {
	Code    int
	Message string
	Success string
}

type APISuccess struct {
	Code    int
	Message string
	Success string
	Data    interface{}
}

// // RevValueBase example
// type RevValueBase struct {
// 	Status bool `json:"Status"`

// 	Err int32 `json:"Err"`
// }

// // RevValue example
// type RevValue struct {
// 	RevValueBase

// 	Data int `json:"Data"`
// }
