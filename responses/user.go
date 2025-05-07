package responses

import "time"

// UserResponse represents a user entity for API responses
type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserLoginResponse represents the response for a successful login
type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

// UsersResponse represents a list of users for API responses
type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Count int            `json:"count"`
}
