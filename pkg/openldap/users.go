package openldap

type User struct {
	Name, Email string
}

func AddUserToGroup(user User, group Group) (*User, error) {
	return nil, nil
}
