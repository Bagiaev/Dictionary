package reports

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) reportExists(id int) (bool, error) {
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM reports WHERE id = $1)`, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

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
	exists, err := r.reportExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("zalupa", id)
	}

	_, err = r.db.Exec(`
		UPDATE reports 
		SET title = $1, description = $2, updated_at = now() 
		WHERE id = $3`, title, description, id)
	return err
}

func (r *Repo) DeleteReport(id int) error {
	exists, err := r.reportExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("zalupa", id)
	}

	_, err = r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	return err
}
