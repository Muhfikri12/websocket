package dashboardrepository

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetEarningDashboard() (int, error)
}

type dashboardRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDashboardRepo(db *gorm.DB, log *zap.Logger) DashboardRepo {

	return &dashboardRepo{db, log}
}

func (dr *dashboardRepo) GetEarningDashboard() (int, error) {

	var totalEarning int
	now := time.Now()

	query := dr.db.Table("orders as o").
		Select("SUM(oi.unit_price * oi.quantity) as total_amount").
		Joins("JOIN order_items as oi ON oi.order_id = o.id").
		Where("o.status = ?", "completed").
		Where("o.created_at BETWEEN ? AND ?",
			StartOfMonth(now),
			EndOfMonth(now),
		).
		Scan(&totalEarning)

	if query.Error != nil {
		return 0, query.Error
	}

	return totalEarning, nil
}

func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// Fungsi untuk mendapatkan akhir bulan
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, -1).Add(24*time.Hour - time.Second)
}
