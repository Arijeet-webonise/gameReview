package models

func (serviceImpl *GenreServiceImpl) GenreByName(_, name string) (*Genre, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM public.genre ` +
		`WHERE name = $1`

	// run query
	XOLog(sqlstr, name)

	g := Genre{}

	err = serviceImpl.DB.QueryRow(sqlstr, name).Scan(&g.ID, &g.Name)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (serviceImpl *GameServiceImpl) FindGameByTitle(_, title string) (*Game, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, developer, summary, rating, image_name, video, video_type ` +
		`FROM public.games ` +
		`WHERE title = $1`

	// run query
	XOLog(sqlstr, title)

	g := Game{}

	err = serviceImpl.DB.QueryRow(sqlstr, title).Scan(&g.ID, &g.Title, &g.Developer, &g.Summary, &g.Rating, &g.ImageName, &g.Video, &g.VideoType)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
