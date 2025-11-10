package application

type UserCreateRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}