package reports

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo { return &Repo{db: db} }

func (r *Repo) GetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repo) CreateReport(title, description string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, now(), now())`, title, description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateReport(id int, title, description string) error {
	_, err := r.db.Exec(`UPDATE reports SET title = $1, description = $2, updated_at = now() WHERE id = $3`, title, description, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
