// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// User represents a row from 'public.users'.
type User struct {
	ID        int            `json:"id"`        // id
	Firstname string         `json:"firstname"` // firstname
	Lastname  sql.NullString `json:"lastname"`  // lastname
	Email     sql.NullString `json:"email"`     // email
	Username  string         `json:"username"`  // username
	Password  string         `json:"password"`  // password
	Roles     sql.NullString `json:"roles"`     // roles

}

type UserService interface {
	DoesUserExists(u *User) (bool, error)
	InsertUser(u *User) error
	UpdateUser(u *User) error
	UpsertUser(u *User) error
	DeleteUser(u *User) error
	GetAllUsers() ([]*User, error)
	GetChunkedUsers(limit int, offset int) ([]*User, error)
}

type UserServiceImpl struct {
	DB XODB
}

// Exists determines if the User exists in the database.
func (serviceImpl *UserServiceImpl) DoesUserExists(u *User) (bool, error) {
	panic("not yet implemented")
}

// Insert inserts the User to the database.
func (serviceImpl *UserServiceImpl) InsertUser(u *User) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.users (` +
		`firstname, lastname, email, username, password, roles` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles)
	err = serviceImpl.DB.QueryRow(sqlstr, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the User in the database.
func (serviceImpl *UserServiceImpl) UpdateUser(u *User) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.users SET (` +
		`firstname, lastname, email, username, password, roles` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE id = $7`

	// run query
	XOLog(sqlstr, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles, u.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles, u.ID)
	return err
}

// Save saves the User to the database.
/*
	func (u *User) Save(db XODB) error {
		if u.Exists() {
			return u.Update(db)
		}

		return u.Insert(db)
	}
*/

// Upsert performs an upsert for User.
//
// NOTE: PostgreSQL 9.5+ only
func (serviceImpl *UserServiceImpl) UpsertUser(u *User) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.users (` +
		`id, firstname, lastname, email, username, password, roles` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, firstname, lastname, email, username, password, roles` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.firstname, EXCLUDED.lastname, EXCLUDED.email, EXCLUDED.username, EXCLUDED.password, EXCLUDED.roles` +
		`)`

	// run query
	XOLog(sqlstr, u.ID, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles)
	_, err = serviceImpl.DB.Exec(sqlstr, u.ID, u.Firstname, u.Lastname, u.Email, u.Username, u.Password, u.Roles)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the User from the database.
func (serviceImpl *UserServiceImpl) DeleteUser(u *User) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.users WHERE id = $1`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers returns all rows from 'public.users',
// ordered by "created_at" in descending order.
func (serviceImpl *UserServiceImpl) GetAllUsers() ([]*User, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.users`

	q, err := serviceImpl.DB.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Username, &u.Password, &u.Roles)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

// GetChunkedUsers returns pagingated rows from 'public.users',
// ordered by "created_at" in descending order.
func (serviceImpl *UserServiceImpl) GetChunkedUsers(limit int, offset int) ([]*User, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.users LIMIT $1 OFFSET $2`

	q, err := serviceImpl.DB.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Username, &u.Password, &u.Roles)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

// UserByID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func (serviceImpl *UserServiceImpl) UserByID(_, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, firstname, lastname, email, username, password, roles ` +
		`FROM public.users ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)

	u := User{}

	err = serviceImpl.DB.QueryRow(sqlstr, id).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Username, &u.Password, &u.Roles)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
