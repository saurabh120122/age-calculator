package models

type CreateUserRequest struct {
	Name string    `json:"name" validate:"required"`
	DOB  DateOnly `json:"dob" validate:"required,pastdate"`
}

type UpdateUserRequest struct {
	Name string    `json:"name" validate:"required"`
	DOB  DateOnly `json:"dob" validate:"required"`
}

type UserResponse struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	DOB  string `json:"dob"`
	Age  int       `json:"age"`
}

type UserBasicResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}
