package model

type User struct {
	Id        int
	Email     string
	Username  string
	FirstName string
	LastName  string
}

type UserAuth struct {
	Id int
}

func (u UserAuth) GetUserAuth() UserAuth {
	return u
}
