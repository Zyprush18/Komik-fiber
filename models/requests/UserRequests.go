package requests

type UserRequests struct{
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email,min=4,max=32"`
	Password string `json:"password" validate:"required"`
}