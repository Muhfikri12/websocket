package dashboardservice

import (
	"project/domain"
	"project/repository"

	"go.uber.org/zap"
)

type DashboardService interface {
	GetEarningProduct() (int, error)
	GetSummary() (*domain.Summary, error)
}

type dashboardService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewDashboardService(repo *repository.Repository, log *zap.Logger) DashboardService {
	return &dashboardService{repo, log}
}

func (ds *dashboardService) GetEarningProduct() (int, error) {

	totalEarning, err := ds.repo.Dashboard.GetEarningDashboard()
	if err != nil {
		return 0, err
	}

	return totalEarning, nil
}

func (ds *dashboardService) GetSummary() (*domain.Summary, error) {

	summary, err := ds.repo.Dashboard.GetSummary()
	if err != nil {
		return nil, err
	}

	return summary, nil
}
