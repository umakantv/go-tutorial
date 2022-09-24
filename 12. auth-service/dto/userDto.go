package dto

type NewUserRequest struct {
	Username   string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	CustomerId string `json:"customer_id"`
	CreatedOn  string `json:"created_on"`
}
