// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// Genretogamerelation represents a row from 'public.genretogamerelation'.
type Genretogamerelation struct {
	ID    int           `json:"id"`    // id
	Game  sql.NullInt64 `json:"game"`  // game
	Genre sql.NullInt64 `json:"genre"` // genre

}

type GenretogamerelationService interface {
	DoesGenretogamerelationExists(g *Genretogamerelation) (bool, error)
	InsertGenretogamerelation(g *Genretogamerelation) error
	UpdateGenretogamerelation(g *Genretogamerelation) error
	UpsertGenretogamerelation(g *Genretogamerelation) error
	DeleteGenretogamerelation(g *Genretogamerelation) error
	GetAllGenretogamerelations() ([]*Genretogamerelation, error)
	GetChunkedGenretogamerelations(limit int, offset int) ([]*Genretogamerelation, error)
}

type GenretogamerelationServiceImpl struct {
	DB XODB
}

// Exists determines if the Genretogamerelation exists in the database.
func (serviceImpl *GenretogamerelationServiceImpl) DoesGenretogamerelationExists(g *Genretogamerelation) (bool, error) {
	panic("not yet implemented")
}

// Insert inserts the Genretogamerelation to the database.
func (serviceImpl *GenretogamerelationServiceImpl) InsertGenretogamerelation(g *Genretogamerelation) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.genretogamerelation (` +
		`game, genre` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, g.Game, g.Genre)
	err = serviceImpl.DB.QueryRow(sqlstr, g.Game, g.Genre).Scan(&g.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the Genretogamerelation in the database.
func (serviceImpl *GenretogamerelationServiceImpl) UpdateGenretogamerelation(g *Genretogamerelation) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.genretogamerelation SET (` +
		`game, genre` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE id = $3`

	// run query
	XOLog(sqlstr, g.Game, g.Genre, g.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, g.Game, g.Genre, g.ID)
	return err
}

// Save saves the Genretogamerelation to the database.
/*
	func (g *Genretogamerelation) Save(db XODB) error {
		if g.Exists() {
			return g.Update(db)
		}

		return g.Insert(db)
	}
*/

// Upsert performs an upsert for Genretogamerelation.
//
// NOTE: PostgreSQL 9.5+ only
func (serviceImpl *GenretogamerelationServiceImpl) UpsertGenretogamerelation(g *Genretogamerelation) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.genretogamerelation (` +
		`id, game, genre` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, game, genre` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.game, EXCLUDED.genre` +
		`)`

	// run query
	XOLog(sqlstr, g.ID, g.Game, g.Genre)
	_, err = serviceImpl.DB.Exec(sqlstr, g.ID, g.Game, g.Genre)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the Genretogamerelation from the database.
func (serviceImpl *GenretogamerelationServiceImpl) DeleteGenretogamerelation(g *Genretogamerelation) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.genretogamerelation WHERE id = $1`

	// run query
	XOLog(sqlstr, g.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, g.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllGenretogamerelations returns all rows from 'public.genretogamerelation',
// ordered by "created_at" in descending order.
func (serviceImpl *GenretogamerelationServiceImpl) GetAllGenretogamerelations() ([]*Genretogamerelation, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.genretogamerelation`

	q, err := serviceImpl.DB.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Genretogamerelation
	for q.Next() {
		g := Genretogamerelation{}

		// scan
		err = q.Scan(&g.ID, &g.Game, &g.Genre)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GetChunkedGenretogamerelations returns pagingated rows from 'public.genretogamerelation',
// ordered by "created_at" in descending order.
func (serviceImpl *GenretogamerelationServiceImpl) GetChunkedGenretogamerelations(limit int, offset int) ([]*Genretogamerelation, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.genretogamerelation LIMIT $1 OFFSET $2`

	q, err := serviceImpl.DB.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Genretogamerelation
	for q.Next() {
		g := Genretogamerelation{}

		// scan
		err = q.Scan(&g.ID, &g.Game, &g.Genre)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// Game returns the Game associated with the Genretogamerelation's Game (game).
//
// Generated from foreign key 'genretogamerelation_game_fkey'.
func (g *Genretogamerelation) GetGames(db XODB) (*Game, error) {
	service := GameServiceImpl{
		DB: db,
	}
	return service.GameByID(int(g.Game.Int64), int(g.Game.Int64))
}

// Genre returns the Genre associated with the Genretogamerelation's Genre (genre).
//
// Generated from foreign key 'genretogamerelation_genre_fkey'.
func (g *Genretogamerelation) GetGenres(db XODB) (*Genre, error) {
	service := GenreServiceImpl{
		DB: db,
	}
	return service.GenreByID(int(g.Genre.Int64), int(g.Genre.Int64))
}

// GenretogamerelationByID retrieves a row from 'public.genretogamerelation' as a Genretogamerelation.
//
// Generated from index 'genretogamerelation_pkey'.
func (serviceImpl *GenretogamerelationServiceImpl) GenretogamerelationByID(_, id int) (*Genretogamerelation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, game, genre ` +
		`FROM public.genretogamerelation ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)

	g := Genretogamerelation{}

	err = serviceImpl.DB.QueryRow(sqlstr, id).Scan(&g.ID, &g.Game, &g.Genre)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
