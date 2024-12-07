package dashboardservice

import (
	"project/repository"

	"go.uber.org/zap"
)

type DashboardService interface {
	GetEarningProduct() (int, error)
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
