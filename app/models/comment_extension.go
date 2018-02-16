package models

func (serviceImpl *CommentServiceImpl) GetGameComments(id int) ([]*Comment, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.comment WHERE game = $1`

	q, err := serviceImpl.DB.Query(sqlstr, id)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Comment
	for q.Next() {
		c := Comment{}

		// scan
		err = q.Scan(&c.ID, &c.Game, &c.Comment, &c.Rating)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

func (serviceImpl *RatingViewServiceImpl) GetGameTotalRating(gameid int) (RatingView, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.rating_view WHERE game = $1`

	// load results
	c := RatingView{}
	err := serviceImpl.DB.QueryRow(sqlstr, gameid).Scan(&c.Game, &c.TotalRating, &c.Count)

	return c, err
}
