package entity

import "demo-app/domain_mydomain/model/errorenum"

type User struct {
	ID       string `` //
	username string
}

type UserRequest struct {
	Username string
}

func NewUser(req UserRequest) (*User, error) {

	if req.Username == "" {
		return nil, errorenum.UsernameMustNotEmpty
	}

	var obj User
	obj.username = req.Username

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *User) Validate() error {
	return nil
}



// assign value here

