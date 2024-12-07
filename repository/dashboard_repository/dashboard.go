package dashboardrepository

import (
	"project/domain"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetEarningDashboard() (int, error)
	GetSummary() (*domain.Summary, error)
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

func (dr *dashboardRepo) GetSummary() (*domain.Summary, error) {
	var summary domain.Summary
	now := time.Now()

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24*time.Hour - time.Second)

	query := dr.db.Table("orders as o").
		Select("SUM(oi.unit_price * oi.quantity) as sales, COUNT(o.id) as orders, SUM(oi.quantity) as items").
		Joins("JOIN order_items as oi ON oi.order_id = o.id").
		Where("o.status != ?", "canceled").
		Where("o.created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Scan(&summary)

	if query.Error != nil {
		return nil, query.Error
	}

	var users int
	queryUser := dr.db.Table("orders as o").
		Select("COUNT(DISTINCT o.customer_id) as users").
		Where("o.created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Scan(&users)

	if queryUser.Error != nil {
		return nil, queryUser.Error
	}

	summary.Users = users

	return &summary, nil
}
