// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// Genre represents a row from 'public.genre'.
type Genre struct {
	ID   int    `json:"id"`   // id
	Name string `json:"name"` // name

}

type GenreService interface {
	DoesGenreExists(g *Genre) (bool, error)
	InsertGenre(g *Genre) error
	UpdateGenre(g *Genre) error
	UpsertGenre(g *Genre) error
	DeleteGenre(g *Genre) error
	GetAllGenres() ([]*Genre, error)
	GetChunkedGenres(limit int, offset int) ([]*Genre, error)
}

type GenreServiceImpl struct {
	DB XODB
}

// Exists determines if the Genre exists in the database.
func (serviceImpl *GenreServiceImpl) DoesGenreExists(g *Genre) (bool, error) {
	panic("not yet implemented")
}

// Insert inserts the Genre to the database.
func (serviceImpl *GenreServiceImpl) InsertGenre(g *Genre) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.genre (` +
		`name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, g.Name)
	err = serviceImpl.DB.QueryRow(sqlstr, g.Name).Scan(&g.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the Genre in the database.
func (serviceImpl *GenreServiceImpl) UpdateGenre(g *Genre) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.genre SET (` +
		`name` +
		`) = ( ` +
		`$1` +
		`) WHERE id = $2`

	// run query
	XOLog(sqlstr, g.Name, g.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, g.Name, g.ID)
	return err
}

// Save saves the Genre to the database.
/*
	func (g *Genre) Save(db XODB) error {
		if g.Exists() {
			return g.Update(db)
		}

		return g.Insert(db)
	}
*/

// Upsert performs an upsert for Genre.
//
// NOTE: PostgreSQL 9.5+ only
func (serviceImpl *GenreServiceImpl) UpsertGenre(g *Genre) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.genre (` +
		`id, name` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, name` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.name` +
		`)`

	// run query
	XOLog(sqlstr, g.ID, g.Name)
	_, err = serviceImpl.DB.Exec(sqlstr, g.ID, g.Name)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the Genre from the database.
func (serviceImpl *GenreServiceImpl) DeleteGenre(g *Genre) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.genre WHERE id = $1`

	// run query
	XOLog(sqlstr, g.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, g.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllGenres returns all rows from 'public.genre',
// ordered by "created_at" in descending order.
func (serviceImpl *GenreServiceImpl) GetAllGenres() ([]*Genre, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.genre`

	q, err := serviceImpl.DB.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Genre
	for q.Next() {
		g := Genre{}

		// scan
		err = q.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GetChunkedGenres returns pagingated rows from 'public.genre',
// ordered by "created_at" in descending order.
func (serviceImpl *GenreServiceImpl) GetChunkedGenres(limit int, offset int) ([]*Genre, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.genre LIMIT $1 OFFSET $2`

	q, err := serviceImpl.DB.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Genre
	for q.Next() {
		g := Genre{}

		// scan
		err = q.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GenreByID retrieves a row from 'public.genre' as a Genre.
//
// Generated from index 'genre_pkey'.
func (serviceImpl *GenreServiceImpl) GenreByID(_, id int) (*Genre, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM public.genre ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)

	g := Genre{}

	err = serviceImpl.DB.QueryRow(sqlstr, id).Scan(&g.ID, &g.Name)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
