package reports

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

//Задание 2
// функция CreateReport
func (r *Repo) RCreateReport(title, description string) error {
	_, err := r.db.Exec(
		`INSERT INTO reports (title, description, created_at, updated_at) 
         VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		title, description,
	)
	if err != nil {
		return err
	}
	return nil
}

// функция для получения репортов GetReport
func (r *Repo) GetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.ID, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

//Функция для обновления UpdateReport
func (r *Repo) RUpdateReport(id int, title, description string) error {
	_, err := r.db.Exec(
		`UPDATE reports 
         SET title = $1, 
             description = $2, 
             updated_at = CURRENT_TIMESTAMP 
         WHERE id = $3`,
		title, description, id,
	)
	if err != nil {
		return err
	}
	return nil
}

//Функция для удаления DeleteReport по id
func (r *Repo) DeleteReportById(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
