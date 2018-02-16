package models

func (serviceImpl *GenretogamerelationServiceImpl)GetGenreOfGame(id int) ([]*Genretogamerelation, error)  {
  const sqlstr = `SELECT ` +
		`*` +
		`FROM public.genretogamerelation WHERE game = $1 `

	q, err := serviceImpl.DB.Query(sqlstr, id)
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
