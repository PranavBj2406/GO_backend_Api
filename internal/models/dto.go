package models

// CreateUserRequest represents a request payload for creating a new user.
// The dob field is expected in YYYY-MM-DD format and will be parsed later in the service layer.
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

// UpdateUserRequest represents a request payload for updating an existing user.
// The dob field is expected in YYYY-MM-DD format and will be parsed later in the service layer.
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

// UserResponse represents the user object returned from API endpoints.
// Age is calculated dynamically in the service layer and is not stored in the database.
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
