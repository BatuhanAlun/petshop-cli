package domain

type IUser interface {
	GetID() int
	GetUsername() string
	CheckPassword(password string) bool
}

type User struct {
	ID        int
	Username  string
	Password  string
	Role      string
	Money     int
	AnimalIDS []int
	ItemIDS   []int
}

type Response struct {
	Data User `json:"Data"`
}
