package repositories

import "database/sql"

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetSummary(start, end string) (int, int, error) {
	var revenue, transaksi int
	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount),0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`, start, end).Scan(&revenue, &transaksi)

	return revenue, transaksi, err
}
