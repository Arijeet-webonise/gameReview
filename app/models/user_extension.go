package models

func (serviceImpl *UserServiceImpl) FindByUsername(username string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, firstname, lastname, username, password, roles ` +
		`FROM public.users ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)

	u := User{}

	err = serviceImpl.DB.QueryRow(sqlstr, username).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Password, &u.Roles)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (serviceImpl *UserServiceImpl) FindByResetLink(resetPassword string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, firstname, lastname, username, password, roles, resetpasswordlink ` +
		`FROM public.users ` +
		`WHERE resetpasswordlink = $1`

	// run query
	XOLog(sqlstr, resetPassword)

	u := User{}

	err = serviceImpl.DB.QueryRow(sqlstr, resetPassword).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Password, &u.Roles, &u.Resetpasswordlink)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
