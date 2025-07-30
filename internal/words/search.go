package words

func (r *Repo) SearchWords(query string) ([]Word, error) {
	var words []Word
	rows, err := r.db.Query(
		`SELECT id, title, translation 
         FROM ru_en 
         ORDER BY similarity(title, $1) DESC
         LIMIT 100`,
		query,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var word Word
		if err := rows.Scan(&word.Id, &word.Title, &word.Translation); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
