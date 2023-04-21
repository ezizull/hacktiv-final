package user

func (n *NewUser) ToDomainMapper() *User {
	return &User{
		UserName: n.UserName,
		Email:    n.Email,
		Age:      n.Age,
		RoleID:   n.RoleID,
	}
}
