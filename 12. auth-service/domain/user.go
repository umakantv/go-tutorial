package domain

type User struct {
	Username   string `json:"email"`
	Password   string `json:""`
	Role       string `json:"role"`
	CustomerId string `json:"customer_id"`
	CreatedOn  string `json:"created_on"`
}

type UserService interface {
	Create(user User) (*User, error)
	Find(username string) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) error
}
