package services

import (
	"kasir-api-golang-2/models"
	"kasir-api-golang-2/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(start, end string) (*models.ReportResponse, error) {
	revenue, transaksi, err := s.repo.GetSummary(start, end)
	if err != nil {
		return nil, err
	}

	return &models.ReportResponse{
		TotalRevenue:   revenue,
		TotalTransaksi: transaksi,
		ProdukTerlaris: models.BestSeller{},
	}, nil
}
